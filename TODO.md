# TODOs

- slog
- json validation, even on config?
- graceful shutdown
- db integration
- a feature that sends multiple concurrent requests
  - uses errgroup
  - uses waitgroup

### CSRF
- https://pkg.go.dev/net/http#CrossOriginProtection
- https://developer.mozilla.org/en-US/docs/Web/Security/Attacks/CSRF#defense_summary_checklist
- set samesite cookies to strict once auth is impl
- no state-changing GETs
