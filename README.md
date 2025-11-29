# Core Platform Monorepo Template

이 레포지토리는 **Core 팀(인프라/공통 로직)과 API 팀(비즈니스 로직) 간의 협업 효율을 극대화하기 위한 Monorepo 템플릿**입니다.

## 🚀 핵심 철학

1.  **Single Source of Truth:** `proto/` 폴더의 Protobuf 파일이 모든 통신의 기준입니다.
2.  **No Packaging:** NPM/Composer 배포 없이, 생성된 코드(`gen/`)를 Git으로 직접 공유합니다.
3.  **Atomic Updates:** 프로토콜 정의, 생성된 코드, 그리고 이를 사용하는 서버 구현체가 하나의 커밋으로 관리됩니다.

## 📂 디렉토리 구조

```text
.
├── buf.work.yaml           # Buf 워크스페이스 설정
├── buf.gen.yaml            # 코드 생성 설정 (Go, TS, PHP, Rust, Python, C++)
├── proto/                  # [Source of Truth] Protobuf 파일
│   ├── buf.yaml
│   └── company/
│       └── auth/v1/
│           └── auth.proto
├── gen/                    # [Generated Code] 생성된 코드가 저장됨 (커밋 포함)
│   ├── go/                 # Go 모듈
│   ├── ts/                 # TS 패키지 (NestJS 호환)
│   ├── php/                # PHP 라이브러리 (Laravel 호환)
│   ├── rust/               # Rust 크레이트 (Serde 지원)
│   ├── python/             # Python 패키지
│   └── cpp/                # C++ 소스
├── sidecar-go/             # [Server] Go 기반 Sidecar (gRPC Server)
├── sidecar-rust/           # [Server] Rust 기반 Sidecar (gRPC Server)
├── tools/                  # [Tools] 빌드 보조 도구
│   └── gen-rust/           # Rust 코드 생성기 (tonic-build)
├── Dockerfile.Go           # Go Sidecar 빌드용 Dockerfile
└── Dockerfile.Rust         # Rust Sidecar 빌드용 Dockerfile
```

## 🛠 Core 팀 가이드 (코드 생성)

Protobuf 파일을 수정한 후, 다음 명령어로 모든 언어의 코드를 생성합니다.

```bash
make generate
```

이 명령어는 내부적으로 다음 작업을 수행합니다:
1.  `buf generate`: Go, TS, PHP, Python, C++ 코드 생성
2.  `cargo run (tools/gen-rust)`: Rust 코드 생성 및 `lib.rs` 자동 갱신 (Serde 지원 포함)

생성된 `gen/` 폴더의 변경 사항을 **반드시 커밋하고 푸시**해야 합니다.

## 📦 API 팀 가이드 (클라이언트 설치)

Git Submodule 대신, 각 언어의 패키지 매니저가 Git 저장소를 직접 참조하도록 설정합니다.

### A. NestJS (TypeScript)
`package.json`에 Git URL을 의존성으로 추가합니다.

```bash
npm install git+https://github.com/my-org/core-platform.git#main
```

### B. Laravel (PHP)
`composer.json`에 레포지토리를 등록하여 설치합니다.

```json
"repositories": [
    {
        "type": "vcs",
        "url": "https://github.com/my-org/core-platform"
    }
],
"require": {
    "my-org/core-proto": "dev-main"
}
```

### C. Python (AI / Data)
`pip`를 사용하여 하위 디렉토리 패키지를 설치합니다.

```bash
pip install git+https://github.com/my-org/core-platform.git#subdirectory=gen/python
```

### D. Go
Go Modules는 하위 디렉토리 모듈을 기본 지원합니다.

```bash
go get github.com/my-org/core-platform/gen/go@latest
```

### E. Rust
Cargo는 Git 저장소의 특정 경로를 지정할 수 있습니다.

```toml
[dependencies]
core-proto = { git = "https://github.com/my-org/core-platform.git", path = "gen/rust" }
```

## 🖥 Sidecar 팀 가이드 (서버 개발)

이 레포지토리 내부에서 개발하므로, `gen/` 폴더를 로컬 경로로 직접 참조합니다.

### Go Sidecar
*   **위치:** `sidecar-go/`
*   **실행:** `go run .`
*   **특징:** `go.mod`의 `replace` 구문을 통해 `../gen/go`를 참조합니다.

### Rust Sidecar
*   **위치:** `sidecar-rust/`
*   **실행:** `cargo run`
*   **특징:** `Cargo.toml`의 `path` 설정을 통해 `../gen/rust`를 참조합니다.

## 🐳 Docker 빌드 및 배포

최적화된 `Dockerfile`이 언어별로 제공됩니다. 빌드 시점에는 코드 생성을 다시 하지 않고, 커밋된 `gen/` 코드를 그대로 사용합니다.

```bash
# Go Sidecar 빌드 및 실행 (Port: 50051)
docker build -f Dockerfile.Go -t sidecar-go .
docker run -p 50051:50051 sidecar-go

# Rust Sidecar 빌드 및 실행 (Port: 50052)
docker build -f Dockerfile.Rust -t sidecar-rust .
docker run -p 50052:50052 sidecar-rust
```

## 📝 라이선스

MPL 2.0 License
