# Module 07: Interfaces and Dynamic Dispatch

## What You'll Learn
Interfaces hold (type, value) pairs. You'll observe dynamic dispatch, see the difference between nil interfaces and interfaces holding nil, and watch type assertions at runtime.

## What to Observe
- Interfaces contain a **runtime type** and a **value**
- The same function can operate on **different types** (dynamic dispatch)
- A **nil interface** ≠ an **interface holding nil**
- Type assertions reveal the **concrete type** at runtime

## Debugging Steps

### Step 1: Interface (Type, Value) Pairs
Set breakpoints at:
1. **Line 48** — `var s Speaker` (nil interface)
2. **Line 54** — After assigning `s = dog`

Start debugging.

At line 48:
- Expand `s` in the Variables panel
- 👀 **`s` is nil** — no type, no value

Press `F5` to reach line 54.
- `s` now holds a `Dog`
- Expand `s` in the debugger
- 👀 **You should see: type=main.Dog, value={name: "Buddy"}**

### Step 2: Dynamic Dispatch
Set breakpoints at:
1. **Line 60** — Before calling `makeItSpeak(dog)`
2. **Line 34** — Inside `makeItSpeak`

Continue to line 60 and press `F11` to step into `makeItSpeak`.

- Look at the parameter `s`
- 👀 **Runtime type is `Dog`**
- Press `F10` to call `s.Speak()`
- The `Dog.Speak` method is called (dynamic dispatch)

Press `Shift+F11` to return to `main`.
- Now call `makeItSpeak(cat)` and step in again
- 👀 **This time, `Cat.Speak` is called** — same code, different behavior

### Step 3: Interface Holding Pointer vs Value
Set a breakpoint at **line 68** (after assigning pointer to interface).

Continue to line 68.
- Expand `s` in the Variables panel
- 👀 **The interface now holds `*Dog` (a pointer type)**

Compare with the earlier assignment (`s = dog`):
- Value: interface holds a **copy** of the struct
- Pointer: interface holds the **address**

### Step 4: Nil Interface vs Interface Holding Nil
Set a breakpoint at **line 77** (after creating both interfaces).

Continue to line 77.
- Expand `nilInterface`: completely nil
- Expand `nonNilInterface`: holds type `*Dog`, but value is `nil`

👀 **Key observation:**
- `nilInterface == nil` is `true`
- `nonNilInterface == nil` is `false` (even though the pointer inside is nil)

This is a common source of bugs. The interface has a type, so it's not nil.

### Step 5: Type Assertions
Set breakpoints at:
1. **Line 85** — Before type assertion
2. **Line 88** — Inside successful assertion
3. **Line 93** — Inside failed assertion

Continue to line 85.
- `s` holds a `Dog`

Press `F10` to execute the type assertion `s.(Dog)`.
- 👀 **Assertion succeeds**, `ok` is `true`
- `d` is the extracted `Dog` value

Press `F5` to reach the failed assertion.
- `s.(Cat)` fails because `s` holds a `Dog`, not a `Cat`
- 👀 **`ok` is `false`**

### Step 6: Type Switch
Set a breakpoint at **line 104** (inside `describeType`).

Continue and step through the type switch.
- 👀 **Watch the `switch` statement determine the runtime type**
- Each case is checked at runtime

## Questions to Answer

1. **What does an interface actually contain?**
   - Is it just a value, or is there more?
   - How does the runtime know which method to call?

2. **Why is an interface holding nil not equal to nil?**
   - What makes an interface "nil"?
   - How can you check if an interface's value is nil?

3. **What's the difference between `Dog` and `*Dog` in an interface?**
   - Does the interface hold a copy or a reference?
   - Can you modify the original struct through the interface?

4. **When do type assertions fail?**
   - What happens if you use the unsafe form `s.(Cat)` on a `Dog`?
   - Try removing the `ok` check and see what happens (panic!)

## Key Takeaway
**Interfaces are (type, value) pairs.** The runtime type determines which method is called. A nil interface has no type. An interface holding nil has a type but a nil value — they are not the same.
