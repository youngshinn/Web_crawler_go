
---


#  Go Web Crawler with Kubernetes Automation

 **주제**: Go 언어로 작성된 뉴스 웹 크롤러를 Kubernetes 환경에서 주기적으로 실행하여 MySQL DB에 저장하는 자동화 시스템 구축

---

## 프로젝트 개요

- **크롤링 대상**: 네이버 뉴스 (검색 키워드 기반)
- **개발 언어**: Go
- **데이터 저장**: MySQL
- **자동화**: Kubernetes CronJob + Argo CD
- **보안 관리**: `.env` 파일을 통한 환경 변수 관리

---

## 주요 기능

1.  **Colly**를 활용한 빠르고 효율적인 웹 크롤링
2.  **GORM**을 이용한 MySQL 데이터 저장
3.  `.env`로 민감한 환경 변수 분리 관리
4.  Docker 기반 애플리케이션 컨테이너화
5.  Kubernetes CronJob을 통한 **주기적 크롤링 자동화**
6.  **GitHub Actions** + **Argo CD** 기반 CI/CD 파이프라인 구축

---

##  CI/CD 파이프라인

| 단계 | 도구 | 설명 |
|------|------|------|
| CI   | **GitHub Actions** | 코드 변경 시 Docker 이미지 자동 빌드 및 레지스트리에 푸시 |
| CD   | **Argo CD**        | Git 저장소의 `k8s/` 매니페스트 변경 감지 후 자동 배포 |

```plaintext
[ GitHub Actions ]
       │
       └── Docker 이미지 빌드 & 푸시
              │
              ▼
[ Argo CD ] ──▶ [ Kubernetes CronJob 자동 배포 ]
```

---

## 프로젝트 구조

```
Goproject/
├── main.go                     # 애플리케이션 진입점
├── go.mod / go.sum             # Go 의존성 관리
├── .env                        # 환경 변수 파일 (Git에 포함되지 않음)
├── Dockerfile                  # Docker 이미지 빌드 설정
├── .github/workflows/
│   └── docker-image.yml        # GitHub Actions CI 설정
├── database/
│   └── database.go             # DB 연결 및 데이터 저장
├── webcrawler/
│   └── crawler.go              # 웹 크롤링 로직
├── k8s/
│   ├── configMap.yaml          # 환경 변수 정의
│   └── cronjob.yaml            # Kubernetes 주기적 실행 정의
```

---

## 크롤링 구조 예시

```go
// 웹 크롤링 후 결과를 DB에 저장할 구조체로 전달
results <- database.News{
	Title:   title,
	Link:    link,
	Keyword: keyword,
}
```

---

## Dockerfile 예시

```dockerfile
FROM golang:1.21
WORKDIR /app
COPY . .
RUN go build -o crawler main.go
CMD ["./crawler", "고속도로", "2"]
```

---

## Kubernetes CronJob 매니페스트 예시

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: crawler-job
spec:
  schedule: "0 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: crawler
            image: your-dockerhub/crawler:latest
            args: ["고속도로", "2"]
            env:
            - name: MYSQL_DSN
              value: "root:password@tcp(mysql-service:3306)/crawlerdb"
          restartPolicy: OnFailure
```


## 주요 기술 스택

| 구분         | 사용 기술                  |
|--------------|----------------------------|
| 언어         | Go                         |
| 데이터베이스 | MySQL                      |
| ORM          | GORM                       |
| 크롤링       | Colly                      |
| 환경 변수    | dotenv                     |
| 컨테이너     | Docker                     |
| 오케스트레이션 | Kubernetes (CronJob)     |
| CI/CD        | GitHub Actions + Argo CD   |

---
---

## 실행 화면
![Image](https://github.com/user-attachments/assets/08004933-455c-4ab8-b763-4ab9b0bda0d4)

![Image](https://github.com/user-attachments/assets/36cdb68c-b29b-49c5-8381-3b5a79fcbf82)

![Image](https://github.com/user-attachments/assets/28f0ecaa-b718-4ee2-ae1d-7dde68f02eda)

---

## 향후 확장 아이디어

- 크롤링 결과를 시각화하는 **대시보드** 구축
- 슬랙/이메일/Discord 등 **알림 연동**
- Prometheus + Grafana 기반 **모니터링 시스템**
- **다중 키워드** 및 **다양한 뉴스 사이트** 크롤링 지원

---
##  작성자 노트

이 프로젝트는 단순한 크롤링을 넘어, **운영 환경까지 고려한 자동화된 백엔드 시스템**입니다.  
웹 크롤링 → DB 저장 → 컨테이너화 → Kubernetes 자동 배포까지 **엔드 투 엔드 전 과정을 직접 구성**한 것이 핵심입니다.

> “단순히 코드를 짜는 것이 아닌, 운영 가능한 구조로 배포까지 책임지는 개발자” 로서의 성장을 보여줄 수 있는 프로젝트입니다.

---
