# Module 01: Main and Entrypoint

## What You'll Learn
Execution does not start at `main()`. You'll observe package initialization, `init()` functions, and program arguments.

## What to Observe
- The `init()` function runs **before** `main()`
- Global variables are initialized before `init()`
- Command-line arguments and environment variables are accessible

## Debugging Steps

### Step 1: Set Breakpoints
Set breakpoints at:
1. **Line 12** — Inside `init()`
2. **Line 18** — Start of `main()`
3. **Line 25** — Where `os.Args` is inspected (before the `fmt.Println` that uses it)
4. **Line 40** — Right before `main()` exits

### Step 2: Start Debugging
- Open "Run and Debug" panel (`Ctrl+Shift+D` / `Cmd+Shift+D`)
- Select **"Debug Module 01 (main-and-entrypoint)"**
- Press `F5`

### Step 3: Watch Execution Flow
When the debugger stops at `init()`:
- Look at the **Call Stack** panel
- Notice `init()` appears before `main()`
- Check `globalCounter` in the **Variables** panel (it should be `0` before line 12 executes)

Press `F10` (Step Over) to execute line 12.
- 👀 **Watch `globalCounter` change to `100`**

### Step 4: Continue to main()
Press `F5` (Continue) to jump to the next breakpoint in `main()`.
- Check the **Call Stack** — now you're in `main.main()`
- 👀 **Look at `globalCounter`** — it's already `100` because `init()` already ran

### Step 5: Inspect Arguments
Press `F5` again to reach the `os.Args` breakpoint (line 25).
- Use the **Watch** panel to add `os.Args` (package-level variables don't always appear in Variables panel)
- Expand `os.Args` to see its contents
- Notice `os.Args[0]` is the program path
- Press `F10` to execute line 26 and see the output

### Step 6: Final Breakpoint
Press `F5` again to reach the final breakpoint at line 40.
- This is right before `main()` finishes
- Check the **Call Stack** — you're still in `main.main()`
- Press `F10` to execute the final `fmt.Println` and watch the program exit

### Step 7: Experiment (Optional)
Stop debugging and configure launch args:
1. Open `.vscode/launch.json`
2. Find the "Debug Module 01" configuration
3. Add:
   ```json
   "args": ["arg1", "arg2", "arg3"]
   ```
4. Debug again and inspect `os.Args`

## Questions to Answer

1. **When does `init()` execute?**
   - Before or after `main()`?
   - Can you call `init()` manually from `main()`?

2. **What is `os.Args[0]`?**
   - Is it the program name or the first argument?

3. **What happens if you have multiple `init()` functions?**
   - Try adding a second `init()` in `main.go` and see the execution order

4. **Can you set a breakpoint BEFORE `init()`?**
   - Where does execution truly start?

## Key Takeaway
**Execution does not begin at `main()`.** Package initialization happens first, including global variable initialization and `init()` functions. When `main()` runs, the world is already set up.
