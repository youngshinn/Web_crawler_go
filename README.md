# Go Web Crawler with Kubernetes Automation

> **주제**: Go 언어로 작성된 뉴스 웹 크롤러를 Kubernetes 환경에서 주기적으로 실행하여 MySQL DB에 저장하는 자동화 시스템 구축

---

## 프로젝트 개요

- **크롤링 대상**: 네이버 뉴스 (검색 키워드 기반)
- **개발 언어**: Go
- **데이터 저장**: MySQL
- **자동화**: Kubernetes CronJob
- **보안 관리**: 환경 변수 `.env` 사용 및 Git Ignore

---

## 주요 기능

1. **Colly 기반 웹 크롤링**
2. **GORM을 활용한 DB 저장**
3. **.env 환경 변수 파일로 DB 연결 정보 관리**
4. **Docker 컨테이너 이미지로 빌드 및 운영**
5. **Kubernetes CronJob을 활용한 주기적 크롤링 자동화**

---

## 크롤링 구조

```go
// 웹 크롤링 후 결과 구조체 채널로 전달
results <- database.News{
	Title:   title,
	Link:    link,
	Keyword: keyword,
}
```

---

## 프로젝트 구조

```
Goproject/
├── main.go
├── go.mod
├── .env              # DB 정보 환경 변수
├── Dockerfile        # 이미지 빌드 설정
├── database/
│   └── database.go   # DB 연결 및 저장 함수
├── webcrawler/
│   └── crawler.go    # 크롤링 로직
```

---



```gitignore
.env
```

---

##  Dockerfile 예시

```dockerfile
FROM golang:1.21
WORKDIR /app
COPY . .
RUN go build -o crawler main.go
CMD ["./crawler", "고속도로", "2"]
```

---

##  Kubernetes CronJob 매니페스트

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

---

##  향후 확장 가능성

- 크롤링 결과에 대한 **시각화 대시보드** 구축
- 슬랙/이메일 알림 연동
- Prometheus + Grafana 로 모니터링 가능
- 다중 키워드 및 뉴스 사이트 확장 지원

---

##  주요 기술 스택

- Go
- GORM
- Colly
- Docker
- Kubernetes (CronJob)
- MySQL
- dotenv

---

##  작성자 노트

이 프로젝트는 백엔드와 클라우드 운영 경험을 모두 보여줄 수 있는 좋은 포트폴리오입니다.  
웹 크롤링 → DB 저장 → 컨테이너화 → 자동화까지 **엔드 투 엔드 전 과정을 직접 구성**한 것이 특징입니다.

> “단순히 코드를 짜는 것이 아닌, 운영 가능한 구조로 배포까지 책임지는 개발자”

으로서의 성장을 보여줄 수 있는 프로젝트입니다. 
