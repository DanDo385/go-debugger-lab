# Context for Claude

## Project Purpose
Educational repository teaching Go debugging with Delve through hands-on examples.
Not a tutorial—a laboratory for building intuition.

## Target Audience
Intermediate Go developers who understand syntax but don't yet "see" runtime behavior in a debugger.

## Design Principles
1. **Show, don't tell**: Every concept must be observable in the debugger
2. **Debug bait**: Intentionally include subtle bugs, shadowing, races
3. **Progressive complexity**: Each module builds on previous mental models
4. **Self-contained**: Each folder is a complete, runnable example
5. **Curiosity-driven**: Ask questions the learner should answer, don't just explain

## Code Style
- Simple, intentional code (not production-style)
- Inline markers: 🔍 breakpoints, 👀 watch variables, ⚠️ surprises
- Comments explain what to OBSERVE, not what code does
- No clever abstractions—this is for learning, not elegance

## When Modifying/Adding Modules
- Ensure VS Code launch configs work for that module
- Include README with: what to observe, where to break, what surprises await
- Test that debugging experience reveals the intended lesson
- Verify it builds with `-gcflags=all=-N -l` (optimizations off)

## Tone
Mentor sitting next to learner, pointing at screen. Precise, curious, respectful of reader's intelligence.
