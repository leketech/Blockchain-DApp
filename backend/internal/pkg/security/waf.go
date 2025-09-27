package security

import (
    "regexp"
    "strings"
)

// WAFConfig holds WAF configuration
type WAFConfig struct {
    BlockSQLInjection    bool
    BlockXSS             bool
    BlockCommandInjection bool
    CustomRules          []string
}

// DefaultWAFConfig returns a default WAF configuration
func DefaultWAFConfig() *WAFConfig {
    return &WAFConfig{
        BlockSQLInjection:    true,
        BlockXSS:             true,
        BlockCommandInjection: true,
        CustomRules:          []string{},
    }
}

// WAF holds WAF implementation
type WAF struct {
    config *WAFConfig
    rules  []*regexp.Regexp
}

// NewWAF creates a new WAF instance
func NewWAF(config *WAFConfig) *WAF {
    if config == nil {
        config = DefaultWAFConfig()
    }

    waf := &WAF{config: config}
    waf.compileRules()
    return waf
}

// compileRules compiles the WAF rules
func (w *WAF) compileRules() {
    // SQL Injection patterns
    sqlPatterns := []string{
        `(?i)(union\s+select)`,
        `(?i)(insert\s+into)`,
        `(?i)(drop\s+table)`,
        `(?i)(delete\s+from)`,
        `(?i)(update\s+\w+\s+set)`,
        `(?i)(select\s+\*\s+from)`,
        `(?i)(or\s+1=1)`,
        `(?i)(or\s+\w+=\w+)`,
        `(?i)(--\s*)`,
        `(?i)(/\*\s*\*/)`,

    }

    // XSS patterns
    xssPatterns := []string{
        `(?i)(<script.*?>)`,
        `(?i)(javascript:)`,
        `(?i)(on\w+\s*=)`,
        `(?i)(eval\s*\()`,
        `(?i)(expression\s*\()`,
        `(?i)(alert\s*\()`,
        `(?i)(document\.cookie)`,

    }

    // Command injection patterns
    cmdPatterns := []string{
        `(?:;|\|\||&|&&)\s*(?:cat|ls|pwd|id|whoami|echo|cp|mv|rm|mkdir|touch|chmod|chown|find|grep|awk|sed|ps|kill|crontab|wget|curl|nc|netcat|telnet|ssh|scp|ftp|tftp|ping|traceroute|dig|nslookup|host|nmap|ncat|nc)\b`,
        `(?:\||&)\s*(?:sh|bash|zsh|ksh|csh|tcsh|fish)\b`,
        `\b(?:exec|system|passthru|shell_exec)\s*\(`,

    }

    // Combine all patterns
    var allPatterns []string
    if w.config.BlockSQLInjection {
        allPatterns = append(allPatterns, sqlPatterns...)
    }
    if w.config.BlockXSS {
        allPatterns = append(allPatterns, xssPatterns...)
    }
    if w.config.BlockCommandInjection {
        allPatterns = append(allPatterns, cmdPatterns...)
    }
    allPatterns = append(allPatterns, w.config.CustomRules...)

    // Compile the patterns
    w.rules = make([]*regexp.Regexp, len(allPatterns))
    for i, pattern := range allPatterns {
        w.rules[i] = regexp.MustCompile(pattern)
    }
}

// CheckRequest checks if a request should be blocked
func (w *WAF) CheckRequest(input string) bool {
    // Normalize the input
    normalized := strings.ToLower(strings.TrimSpace(input))

    // Check against all rules
    for _, rule := range w.rules {
        if rule.MatchString(normalized) {
            return true // Block the request
        }
    }

    return false // Allow the request
}