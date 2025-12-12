# Development Workflow & Release Strategy

This document defines the standard operating procedure for the `unik` project, following strict Gitflow principles and Semantic Versioning.

## 1. Branching Strategy

We use a 3-tier branching model:

### 🟢 `dev` (Development)

- **Purpose**: The "Bleeding Edge". All new features and non-critical bug fixes start here.
- **Workflow**:
  1.  Create a feature branch from `dev`: `git checkout -b feat/my-feature dev`.
  2.  Work, Commit (using Conventional Commits), Push.
  3.  Open Pull Request (PR) -> `dev`.
  4.  CI runs Tests. Merge upon success.
- **Stability**: Low. May contain experimental code.

### 🟡 `staging` (Pre-Release / Beta)

- **Purpose**: Stabilization and QA before public release.
- **Workflow**:
  1.  When ready to prepare a release, merge `dev` -> `staging` via PR.
  2.  No new features are added directly here. Only bug fixes found during testing (`fix:` commits).
  3.  This triggers "Release Candidate" builds if configured (e.g., `v0.2.0-rc1`).

### 🔴 `main` (Production)

- **Purpose**: Stable, tagged releases.
- **Workflow**:
  1.  When `staging` is stable, create a PR: `staging` -> `main`.
  2.  **Release Please** Action runs on `main`. It analyzes commits since the last release.
  3.  It creates a "Release PR".
  4.  Merging the "Release PR" creates the Git Tag (e.g., `v1.0.0`) and GitHub Release.

---

## 2. Versioning Criteria (SemVer)

We adhere to [Semantic Versioning](https://semver.org/).

### 🚀 Major Version (v1.0.0, v2.0.0)

A major release signifies a **stable, production-ready milestone** or **breaking changes**.

**Checklist for `v1.0.0` (Definition of Done):**

- [ ] **Stable API**: No planned changes to public function signatures (`Smart`, `Regional`, `Duration`).
- [ ] **Deep Localization**: Support for major Asian languages (ID, TH, VN, JP, CN) is complete and tested.
- [ ] **I18n Engine**: Robust pluralization support (not just appending "s").
- [ ] **Test Coverage**: Minimum 80% coverage on core packages.
- [ ] **Documentation**: Complete GoDocs and README for all public exports.

**When to bump Major (v2+):**

- Removing a public function.
- Changing a function signature (e.g., adding a required parameter).
- Changing behavior significantly (e.g., `Smart` returning different logic that breaks implementation).

### ✨ Minor Version (v0.x.0, v1.x.0)

New features that are **backward-compatible**.

**Examples for `unik`:**

- Adding a new Region constant (e.g., `RegionBR` for Brazil).
- Adding a new utility function (e.g., `func DurationShort(...)`).
- Adding support for a new Language (without changing existing API).
- Performance improvements (Internal refactors).

### 🐛 Patch Version (v0.x.x, v1.x.x)

Backward-compatible bug fixes.

**Examples for `unik`:**

- Fixing a typo in a translation ("Just nwo" -> "Just now").
- Fixing a timezone calculation bug.
- Fixing a crash in `Parse`.

---

## 3. Workflow Summary for Developer

| Action          | Branch    | Command/Commit                | Note                      |
| :-------------- | :-------- | :---------------------------- | :------------------------ |
| Start Feature   | `dev`     | `git checkout -b feat/x dev`  | Always start from dev     |
| Commit Work     | Feature   | `git commit -m "feat: add X"` | Use conventional commits! |
| Merge Feature   | `dev`     | Pull Request                  | Squash merge              |
| Prepare Release | `staging` | PR `dev` -> `staging`         | QA happens here           |
| **Release**     | `main`    | PR `staging` -> `main`        | Triggers Release Please   |
