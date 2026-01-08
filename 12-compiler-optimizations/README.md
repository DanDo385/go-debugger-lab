# Module 12: Compiler Optimizations

## What You'll Learn
Optimizations make debugging hard. You'll observe variables disappearing, breakpoints failing, and understand why production binaries lie to debuggers.

## What to Observe
- **Inlining** makes function calls disappear
- **Dead code elimination** removes "unused" variables
- **Register allocation** means variables have no memory address
- Variables show **`<optimized out>`** in the debugger

## Debugging Steps

### Step 1: Optimizations OFF (Default)
This is the default configuration in this lab.

Set breakpoints at:
1. **Line 7** — Inside `add` function
2. **Line 14** — Inside `calculate` function
3. **Line 16** — On `temp1`

Debug normally:
- All breakpoints trigger
- All variables are visible
- Stepping works predictably

### Step 2: Optimizations ON (Experiment)
To see the effect of optimizations, you need to build an optimized binary.

**Option 1: Change launch config temporarily**
1. Open `.vscode/launch.json`
2. Find "Debug Module 12"
3. Change `"buildFlags": "-gcflags=all=-N -l"` to `"buildFlags": ""`
4. Debug again

**Option 2: Use command-line Delve**
```bash
cd 12-compiler-optimizations
go build -o optimized
dlv exec ./optimized
```

Set breakpoints:
```
(dlv) break main.add
(dlv) break main.calculate
(dlv) continue
```

👀 **Observe:**
- Breakpoint in `add` may not trigger (function was inlined)
- Variables in `calculate` show `<optimized out>`
- Stepping may jump over multiple lines

### Step 3: Compare Side by Side
Run the same program in both modes:

**Unoptimized:**
```bash
go build -gcflags=all=-N -l -o unoptimized
dlv exec ./unoptimized
```

**Optimized:**
```bash
go build -o optimized
dlv exec ./optimized
```

In each debugger, try:
- `break main.add`
- `continue`
- `print temp1` (in `calculate`)

## Questions to Answer

1. **What does "inlining" mean?**
   - Why does the `add` breakpoint not trigger?
   - Where did the function call go?

2. **Why do variables show `<optimized out>`?**
   - Does the variable still exist?
   - Where is the value stored?

3. **Why does Go disable optimizations for debugging?**
   - What's the trade-off?
   - Why not always disable optimizations?

4. **How does this affect production debugging?**
   - Can you reliably debug an optimized production binary?
   - What alternative strategies exist? (Logging, profiling, distributed tracing)

5. **Try debugging a release build of a real Go program:**
   - Find a Go binary (e.g., `which go`)
   - Try: `dlv exec $(which go)`
   - Set a breakpoint and see what happens

## Key Takeaway
**Optimizations lie to debuggers.** Inlining, dead code elimination, and register allocation make variables and function calls disappear. Always debug with `-gcflags=all=-N -l` (optimizations off). Production debugging requires different tools: logging, profiling, distributed tracing.
