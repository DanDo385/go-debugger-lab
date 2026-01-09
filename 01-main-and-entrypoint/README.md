# Module 01: Main and Entrypoint

## What You'll Learn
Execution does not start at `main()`. You'll observe package initialization, `init()` functions, and program arguments.

## What to Observe
- The `init()` function runs **before** `main()`
- Global variables are initialized before `init()`
- Command-line arguments and environment variables are accessible

## Debugging Steps

### Step 1: Set Breakpoints
**⚠️ CRITICAL: Set breakpoints BEFORE starting the debugger!**

If you start debugging without breakpoints, the program will run to completion and exit (you'll see "Process exited with status 0" — that's normal, but not what we want).

**Important:** Set breakpoints on **executable lines** (code), not comment lines. VS Code may not stop reliably on comment-only lines.

Set breakpoints at:
1. **Line 12** — Click in the left margin (gutter) next to `func init()` — a red dot should appear
2. **Line 18** — Click next to `func main()` — another red dot
3. **Line 26** — Click next to `fmt.Println("Program name:", os.Args[0])` ⚠️ **NOT line 25** (the comment)
4. **Line 41** — Click next to `fmt.Println("main() finished")` ⚠️ **NOT line 40** (the comment)

**Verify:** You should see 4 red dots in the left margin before proceeding. Make sure they're on lines 12, 18, 26, and 41 (the executable lines).

### Step 2: Start Debugging

**⚠️ IMPORTANT:** Don't just press `F5` directly on the file! You need to select the correct launch configuration.

**Method 1: Using Run and Debug Panel (Recommended)**
1. Open "Run and Debug" panel (`Ctrl+Shift+D` / `Cmd+Shift+D`)
2. In the dropdown at the top, select **"Debug Module 01 (main-and-entrypoint)"**
3. Press `F5` or click the green play button

**Method 2: Quick Debug**
1. Open `main.go`
2. Press `F5`
3. When prompted, select **"Debug Module 01 (main-and-entrypoint)"** from the dropdown
   - ⚠️ **Don't select "Debug Current File"** — that won't work correctly

**If the program runs without stopping:**
- ❌ **Problem:** You didn't set breakpoints before starting
- ✅ **Fix:** Stop debugging, set breakpoints (red dots), then start again
- The debugger will stop at the first breakpoint (line 12 in `init()`)

**If breakpoints are set but F5 (Continue) skips them:**
- ❌ **Problem:** Breakpoints are on comment lines instead of executable lines
- ✅ **Fix:** 
  1. Remove breakpoints on lines 25 and 40 (comment lines)
  2. Set breakpoints on lines 26 and 41 (executable lines)
  3. Red dots should be on: 12, 18, **26**, **41** (not 25 or 40)
- VS Code debugger needs executable code to stop reliably

**If it crashes or fails:**
1. Check that Delve is installed: `dlv version`
2. Check that Go tools are installed (VS Code should prompt you)
3. Try cleaning the debug binary: Delete `__debug_bin*` files in this directory
4. Restart VS Code
5. Make sure you're in the workspace root, not just the module folder

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
Press `F5` again to reach the `os.Args` breakpoint (line 26).
- ⚠️ **If it doesn't stop:** Check that your breakpoint is on line 26 (executable), not line 25 (comment)
- Use the **Watch** panel to add `os.Args` (package-level variables don't always appear in Variables panel)
- Expand `os.Args` to see its contents
- Notice `os.Args[0]` is the program path
- Press `F10` to execute line 26 and see the output

### Step 6: Final Breakpoint
Press `F5` again to reach the final breakpoint at line 41.
- ⚠️ **If it doesn't stop:** Check that your breakpoint is on line 41 (executable), not line 40 (comment)
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

## Troubleshooting

### "It crashed when I pressed F5"

**Common causes and fixes:**

1. **Wrong launch configuration selected**
   - ❌ Don't use "Debug Current File" 
   - ✅ Use "Debug Module 01 (main-and-entrypoint)"
   - Fix: Always select from the dropdown in Run and Debug panel

2. **Delve not installed**
   - Error: `dlv: command not found`
   - Fix: `go install github.com/go-delve/delve/cmd/dlv@latest`

3. **Go tools not installed**
   - VS Code should prompt you to install Go tools
   - Fix: Click "Install All" when prompted, or run: `go install -a github.com/go-delve/delve/cmd/dlv@latest`

4. **Corrupted debug binary**
   - Error: Permission denied or binary won't run
   - Fix: Delete `__debug_bin*` files in this directory and try again

5. **Wrong working directory**
   - Make sure you opened the workspace root (`go-debugger-lab`), not just the module folder
   - Fix: File → Open Folder → select the `go-debugger-lab` folder

6. **Go version mismatch**
   - Check `go.mod` requires Go 1.25
   - Fix: Update Go or adjust `go.mod` if needed

**Still having issues?**
- Check the Debug Console (bottom panel) for error messages
- Check the Output panel → select "Go" from dropdown
- Try running manually: `cd 01-main-and-entrypoint && go run main.go`

## Key Takeaway
**Execution does not begin at `main()`.** Package initialization happens first, including global variable initialization and `init()` functions. When `main()` runs, the world is already set up.
