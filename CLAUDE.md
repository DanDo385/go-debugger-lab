# CLAUDE.md — Maintenance Context for AI Assistants

## Project Purpose

This is an **educational repository** teaching Go debugging with Delve through hands-on examples.

**Not a tutorial — a laboratory for building intuition.**

The goal is to make invisible runtime behavior visible. Every module is designed to answer questions like:
- "Where is this value stored?"
- "Why did this variable change?"
- "What is the goroutine actually doing?"

## Target Audience

**Intermediate Go developers** who:
- Understand syntax: functions, structs, pointers, goroutines, channels, tests
- Do NOT yet "see" runtime behavior in their mind
- Cannot predict stack frames, heap escapes, or goroutine interleaving
- Have used a debugger but only to "step randomly until something looks wrong"

After completing this lab, they should be able to **observe execution with precision** and **trust (or question) the debugger intelligently**.

---

## Design Principles

### 1. Show, Don't Tell
Every concept must be **observable in the debugger**.

Bad:
> "Pointers store addresses."

Good:
> "Set a breakpoint at line 12. Watch the address of `x`. Now step into the function. Does the address change? Why?"

### 2. Debug Bait
Intentionally include:
- **Shadowed variables** (same name, different scope)
- **Subtle bugs** (off-by-one, race conditions)
- **Surprising behavior** (deferred functions, interface nil vs nil interface)

The code is simple, but the *observations* are interesting.

### 3. Progressive Complexity
Each module builds on previous mental models:
- Module 01: "Execution starts at `main()`"
- Module 02: "Variables exist in scopes"
- Module 03: "Each function call creates a stack frame"
- Module 09: "Goroutines make stepping feel broken"

Do not skip ahead to concurrency without establishing stack frames first.

### 4. Self-Contained Modules
Each folder is a complete, runnable example.
- `main.go` (or `*_test.go`)
- `README.md` with debugging instructions
- No shared dependencies between modules

### 5. Curiosity-Driven
Ask questions the learner should answer. Do not just explain.

Bad:
> "This variable is shadowed because Go allows redeclaration in inner scopes."

Good:
> "Set a breakpoint at line 8 and line 12. Are `x` and `x` the same variable? Check their addresses."

---

## Code Style

### Keep It Simple
This is for **learning**, not production.

Avoid:
- Clever abstractions
- Error handling beyond what's necessary
- "Best practices" that obscure the lesson

Prefer:
- Flat, linear code
- Obvious names (`original`, `copy`, `escaped`)
- Intentional simplicity

### Inline Markers
Use emoji markers to guide the learner:
```go
// 🔍 SET BREAKPOINT HERE
// 👀 WATCH THIS VARIABLE
// ⚠️ WHY DOES THIS CHANGE?
// 🤔 DOES THIS ESCAPE TO THE HEAP?
```

### Comments Explain What to Observe
Not what the code does.

Bad:
```go
// Increment i
i++
```

Good:
```go
i++ // 👀 Watch i in the Variables panel
```

---

## When Modifying or Adding Modules

### Checklist
- [ ] Ensure the VS Code launch config works for this module
- [ ] Include `README.md` with:
  - What to observe
  - Where to set breakpoints
  - What surprises to expect
  - Questions to answer
- [ ] Test that the debugging experience reveals the intended lesson
- [ ] Verify it builds with `-gcflags=all=-N -l` (optimizations off)
- [ ] Ensure the code is intentionally simple (no unnecessary complexity)

### README.md Template for Modules
```markdown
# Module XX: [Topic Name]

## What You'll Learn
[One sentence: what runtime behavior will you observe?]

## What to Observe
- [Specific thing 1]
- [Specific thing 2]
- [Specific thing 3]

## Debugging Steps

### Step 1: [Action]
- Set breakpoint at [location]
- [What to look for in Variables panel / Call Stack / etc.]

### Step 2: [Action]
- Step Over / Step Into / Continue
- [What changed? What stayed the same?]

### Step 3: [Action]
- [Final observation]

## Questions to Answer
1. [Question that requires observing runtime behavior]
2. [Question that challenges assumptions]
3. [Question that builds mental model]

## Key Takeaway
[One sentence summarizing the mental model this module builds]
```

---

## Tone and Voice

Write like a **mentor sitting next to the learner, pointing at the screen**.

### Good Examples
- "Watch what happens when you step into this function."
- "Did you expect that address to change?"
- "Notice the call stack now has three frames."

### Bad Examples
- "As we all know, Go uses pass-by-value semantics..."
- "This is a common mistake that beginners make..."
- "Let's leverage the power of the debugger to..."

### Principles
- **Precise, not verbose**: Use short sentences.
- **Curious, not authoritative**: Ask questions, don't lecture.
- **Respectful**: Assume the reader is intelligent, just learning to observe.

---

## Build and Debug Configuration

### Always Disable Optimizations
All modules must build with:
```bash
go build -gcflags=all=-N -l
```

This is configured in `.vscode/launch.json` automatically.

### Why?
Optimized builds:
- Inline functions (breaks stepping)
- Eliminate "unused" variables (hides state)
- Store variables in registers (no address to inspect)

Module 12 (compiler-optimizations) demonstrates this explicitly by comparing optimized vs unoptimized builds.

---

## Module-Specific Notes

### Module 01: main-and-entrypoint
- Show `init()` runs before `main()`
- Demonstrate how execution flows through package initialization

### Module 02: variables-and-scope
- Use shadowing intentionally
- Show loop variable capture (Go 1.22 vs older versions if relevant)

### Module 03: functions-and-call-stack
- Include recursion
- Show stack frames building up and tearing down

### Module 04: pointers-and-memory
- Print memory addresses explicitly
- Show pass-by-value vs pointer mutation

### Module 05: slices-maps-and-aliasing
- Demonstrate shared backing arrays
- Show how modifying one slice affects another

### Module 06: structs-and-methods
- Value receiver vs pointer receiver
- Show when methods mutate vs when they don't

### Module 07: interfaces-and-dynamic-dispatch
- Show `(type, value)` pairs in debugger
- Demonstrate `nil` interface vs interface holding `nil` pointer

### Module 08: errors-and-defer
- Show deferred functions execute AFTER return statement
- Demonstrate named returns + defer

### Module 09: goroutines-basics
- Show how stepping "skips" to different goroutines
- Demonstrate goroutine panel in debugger

### Module 10: channels-and-blocking
- Show blocked goroutines
- Demonstrate unbuffered vs buffered channels

### Module 11: data-races-and-sync
- **Critical**: Show how the debugger changes race behavior (Heisenbug)
- Compare: no debugger, with debugger, with `-race` flag
- Teach when to trust debugger vs when it lies

### Module 12: compiler-optimizations
- Compare same code with and without optimizations
- Show variables disappearing
- Demonstrate why production debugging is hard

### Module 13: debugging-tests
- Debug a failing test
- Show table-driven tests
- Demonstrate conditional breakpoints in test loops

---

## Common Pitfalls to Avoid

### Don't Over-Engineer
The code should be simple enough to fit in your head while debugging.

### Don't Skip the README
Every module needs clear instructions. The learner should not have to guess where to set breakpoints.

### Don't Assume Prior Knowledge of Debuggers
Even if they've used a debugger before, they may not know:
- What the call stack represents
- How to watch expressions
- How to view goroutines

### Don't Make It a Reference Manual
This is a lab, not documentation. Focus on **observation and discovery**, not comprehensive coverage.

---

## Testing Your Changes

Before committing a new or modified module:

1. **Build it:**
   ```bash
   cd XX-module-name
   go build -gcflags=all=-N -l
   ```

2. **Debug it:**
   - Open the module in VS Code
   - Select the corresponding launch config
   - Verify breakpoints trigger
   - Verify the intended behavior is observable

3. **Read the README as a new learner:**
   - Can you follow the instructions without prior context?
   - Do the questions make sense?
   - Is the payoff clear?

---

## Philosophy

A debugger is not a crutch. It's a **microscope**.

This repository teaches developers to:
- **Observe** runtime behavior precisely
- **Question** their assumptions
- **Trust** the debugger (and know when not to)

Every module should leave the learner thinking:
> "I didn't know I could see that."

---

## Maintenance Guidelines

### When Adding a New Module
1. Follow the design principles above
2. Add a launch config in `.vscode/launch.json`
3. Update the module table in root `README.md`
4. Ensure it builds and debugs correctly
5. Test the README instructions with fresh eyes

### When Modifying Existing Modules
1. Preserve the original lesson
2. Keep the code simple
3. Test that the debugging experience still works
4. Update the README if the steps change

### When Fixing Bugs
- If the bug is intentional (debug bait), document it in the README
- If the bug is unintentional, fix it without adding complexity

---

## Final Notes

This repository is a **teaching tool**, not a production codebase.

Optimize for:
- **Clarity of observation**
- **Ease of debugging**
- **Aha moments**

Do not optimize for:
- Code elegance
- Performance
- "Best practices" that obscure the lesson

The measure of success is:
> "After this module, can the learner answer the question without guessing?"

If yes, the module works.
