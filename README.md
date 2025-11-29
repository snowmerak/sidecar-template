# Core Platform Monorepo Template

이 레포지토리는 **Core 팀(인프라/공통 로직)과 API 팀(비즈니스 로직) 간의 협업 효율을 극대화하기 위한 Monorepo 템플릿**입니다.

## 🚀 핵심 철학

1.  **Single Source of Truth:** `proto/` 폴더의 Protobuf 파일이 모든 통신의 기준입니다.
2.  **No Packaging:** NPM/Composer 배포 없이, 생성된 코드(`gen/`)를 Git으로 직접 공유합니다.
3.  **Atomic Updates:** 프로토콜 정의와 서버 구현체가 하나의 커밋으로 관리됩니다.

## 📂 디렉토리 구조

```text
.
├── buf.yaml                # Buf 모듈 정의
├── buf.gen.yaml            # 코드 생성 설정 (Go, TS, PHP, Rust 등)
├── proto/                  # [Source of Truth] Protobuf 파일
│   └── company/
│       └── auth/v1/
│           └── auth.proto
├── gen/                    # [Generated Code] 생성된 코드가 저장됨 (커밋 포함)
│   ├── go/                 # Go 모듈 (Sidecar가 사용)
│   ├── ts/                 # TS 패키지 (NestJS가 사용)
│   ├── php/                # PHP 클래스 (Laravel이 사용)
│   └── rust/               # Rust 크레이트 (Sidecar가 사용)
├── sidecar-go/             # [Server] Go 기반 Sidecar
│   ├── main.go
│   └── go.mod              # gen/go를 로컬 경로로 참조 (replace)
└── sidecar-rust/           # [Server] Rust 기반 Sidecar
    ├── src/main.rs
    └── Cargo.toml          # gen/rust를 로컬 경로로 참조 (path)
```

## 🛠 사용 방법

### 1. 초기 설정

이 템플릿을 클론한 후, 프로젝트에 맞게 `buf.gen.yaml`과 `go.mod` 등의 패키지명을 수정하세요.

*   `buf.gen.yaml`: `go_package_prefix` 수정
*   `gen/go/go.mod`: 모듈명 수정
*   `sidecar-go/go.mod`: `require` 및 `replace` 경로 수정

### 2. 코드 생성 (Core 팀)

Protobuf 파일을 수정한 후, 다음 명령어로 코드를 생성합니다.

```bash
buf generate
```

생성된 `gen/` 폴더의 변경 사항을 커밋하고 푸시합니다.

### 3. API 팀 연동 (NestJS / Laravel)

이 레포지토리를 **Git Submodule**로 추가하여 사용합니다.

```bash
# 예시: libs/core 경로에 서브모듈 추가
git submodule add <repo-url> libs/core
```

*   **NestJS:** `tsconfig.json`의 `paths`에 `@core/*`: `["libs/core/gen/ts/*"]` 추가
*   **Laravel:** `composer.json`의 `autoload`에 네임스페이스 매핑 추가

### 4. Sidecar 개발 (Go / Rust)

`sidecar-go` 또는 `sidecar-rust` 디렉토리에서 개발을 진행합니다.
이미 로컬 경로로 `gen/` 폴더가 연결되어 있으므로, 별도의 설치 없이 생성된 코드를 즉시 사용할 수 있습니다.

## 📝 라이선스

MIT License
