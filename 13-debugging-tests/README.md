# Module 13: Debugging Tests

## What You'll Learn
Tests are just code. You'll debug failing tests, set conditional breakpoints in table-driven tests, and understand test helpers.

## What to Observe
- **Tests are debuggable** like any other code
- **Table-driven tests** let you debug specific cases
- **Conditional breakpoints** target specific test rows
- **`t.Helper()`** improves error reporting

## Debugging Steps

### Step 1: Debug a Table-Driven Test
Set breakpoints at:
1. **Line 21** — Inside `t.Run` (where `Add` is called)
2. **Line 26** — Before the assertion

**Option A:** Use VS Code launch config
- Select **"Debug Module 13 (debugging-tests)"**
- Press `F5`

**Option B:** Use the testing UI
- In VS Code, you'll see "run | debug" above each test function
- Click **"debug"** next to `TestAdd`

When the debugger stops:
- Expand `tt` in the Variables panel
- 👀 See `tt.name`, `tt.a`, `tt.b`, `tt.expected`

Press `F5` to continue through each test case.

### Step 2: Conditional Breakpoint
Right-click on the breakpoint at line 21.
- Select **"Edit Breakpoint"**
- Add condition: `tt.name == "negative numbers"`
- Press `F5`

👀 **The debugger only stops for that specific test case.**

### Step 3: Debug a Failing Test
Run the tests first to see which fail:
```bash
cd 13-debugging-tests
go test -v
```

Notice `TestDivide/division_with_remainder` passes (because `10/3 = 3` in integer division).

Set a breakpoint at **line 59** (inside `division with remainder` subtest).

Debug the test:
- Step through `Divide(10, 3)`
- 👀 Watch `result` — it's `3` (integer division)
- The test expects `3`, so it passes
- But conceptually, this might be surprising

### Step 4: Debug FindMax
Set a breakpoint at **line 90** (inside `TestFindMax`).

Debug the test and step into `FindMax` for the "multiple values" case.

Inside `FindMax`:
- The loop starts at `i := 0`
- 👀 **This compares `nums[0]` with itself** — unnecessary but harmless

Try the "single value" case:
- `max` starts at `nums[0]`
- The loop still runs and compares `nums[0]` with itself
- The bug is inefficient but doesn't cause a test failure

**This is a subtle bug** — the loop should start at `i := 1`.

### Step 5: Test Helper
Set breakpoints at:
1. **Line 115** — Calling `assertEqual`
2. **Line 109** — Inside `assertEqual`

Debug `TestWithHelper`:
- Step into `assertEqual`
- Notice `t.Helper()` at the top
- This marks the function as a test helper

If `assertEqual` fails:
- The error points to **the caller** (line 115), not the helper (line 109)
- This makes test failures clearer

### Step 6: Running a Specific Test
From the terminal:
```bash
# Run only TestAdd
go test -run TestAdd -v

# Run only the "negative numbers" subtest
go test -run TestAdd/negative -v

# Debug a specific test
dlv test -- -test.run TestAdd
```

## Questions to Answer

1. **How do you debug a specific test case in a table-driven test?**
   - Use conditional breakpoints?
   - Run a specific subtest?

2. **Why use `t.Helper()`?**
   - Try removing it from `assertEqual` — what changes in the error message?

3. **Can you debug benchmarks?**
   - Set a breakpoint in `BenchmarkAdd`
   - Debug the test — what happens?

4. **How do you find which test is failing?**
   - Run `go test -v` first
   - Set breakpoints in the failing test
   - Debug and observe

5. **Try fixing the bug in `FindMax`:**
   - Change `for i := 0` to `for i := 1`
   - Do the tests still pass?
   - Is the function more efficient?

## Key Takeaway
**Tests are debuggable.** Use conditional breakpoints for table-driven tests. Use `t.Helper()` in test helpers for clearer errors. Debug failing tests to observe the actual vs expected behavior.
