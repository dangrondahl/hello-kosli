package version

// GitSHA is injected via -ldflags at build time.
// Defaults to "dev" for local runs and tests.
var GitSHA = "dev"
