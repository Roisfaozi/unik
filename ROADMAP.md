# Unik Package Development Roadmap

This document outlines the strategic development plan for `unik`. The goal is to evolve from a "Format Utility" to a "Comprehensive Time Engine" for Go applications.

## Phase 1: Deep Localization & Internationalization (Currently In Progress)

**Objective**: Remove "English-centric" limitations in relative time and formatting.

### 1.1 Multi-Language Dictionary Engine (✅ Completed)

- [x] Refactor `social.go` to use a `Locale` based dictionary.
- [x] Implement robust PluralRules (e.g., "1 minute" vs "2 minutes").
- [x] Basic support for EN and ID.

### 1.2 Expanded Language Support (✅ Completed)

- [x] **Thailand (TH)**: Add dictionary and specific Particles.
- [x] **Vietnam (VN)**: Add dictionary.
- [x] **Japan (JP)**: Add dictionary (Counter suffix support).
- [x] **Malaysia (MY)**: Add dictionary (similar to ID but distinct).
- [ ] **Arabic (AR)**: (Optional) Complex dual plural handling.

### 1.3 Native Era Support (✅ Completed)

- [x] Create `CalendarSystem` interface.
- [x] **Japanese Era**: Support changing `2024` -> `Reiwa 6`.
- [x] **Hijri Calendar**: Support Islamic date conversion (Tabular).
- [x] Update `Regional` function to accepted `WithCalendar(...)` option.

## Phase 2: Robust & Fuzzy Parsing (v0.3.0)

**Objective**: Make input handling more forgiving and intelligent.

### 2.1 Fuzzy Parsing Engine

- [ ] Implement "Best Effort" parsing.
- [ ] Detect delimiters automaticallly (`/`, `-`, `.`, ` `).
- [ ] **Heuristic Logic**:
  - If `part[0] > 12`, it must be Day (valid for US vs EU ambiguity resolution attempt).
  - If `part[0]` is 4 digits, it's YYYY.

### 2.2 Strict Configuration

- [ ] Add `WithStrict(bool)` option.

## Phase 3: Performance & Ecosystem (v0.4.0)

**Objective**: Ensure suitability for high-throughput systems.

### 3.1 Allocations Benchmark

- [ ] Create `benchmark_test.go`.
- [ ] Target: Zero allocation for `Format` calls if possible (reuse buffers).

### 3.2 Framework Midddleware

- [ ] **Gin/Echo Middleware**: Auto-detect `Accept-Language` header.

### 3.3 CLI Tooling

- [ ] Build a standalone binary `unik`.
