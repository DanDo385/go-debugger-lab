# Module 10: Channels and Blocking

## What You'll Learn
Channels cause goroutines to block. You'll observe unbuffered vs buffered channels, see blocked goroutines in the debugger, and understand select statements.

## What to Observe
- **Unbuffered channels** block until both sender and receiver are ready
- **Buffered channels** don't block until the buffer is full
- Blocked goroutines appear in the **Goroutines panel**
- `select` waits for **the first available channel**
- Receiving from a **closed channel** returns the zero value

## Debugging Steps

### Step 1: Unbuffered Channel Blocking
Set breakpoints at:
1. **Line 35** — Before launching receiver
2. **Line 22** — Inside receiver (waiting)
3. **Line 42** — Before sending

Continue to line 35 and launch the receiver.
- Check the **Goroutines panel**
- 👀 The receiver goroutine is **blocked** on `<-ch`

Press `F5` to reach line 42.
- Main is about to send

Press `F10` to send `42`.
- 👀 **The receiver unblocks immediately**
- Unbuffered channels are synchronous

### Step 2: Buffered Channel
Set breakpoints at:
1. **Line 53** — First send
2. **Line 62** — First receive

Continue and watch the sends.
- Buffer has capacity 2
- 👀 **Sends don't block** until buffer is full

Press `F5` to reach the receives.
- Receives drain the buffer

### Step 3: Select Statement
Set a breakpoint at **line 97** (inside `select`).

Continue and watch which case executes.
- 👀 **`ch1` case executes** (shorter delay)
- `select` waits for the first ready channel

### Step 4: Closed Channels
Set a breakpoint at **line 125** (receiving from closed channel).

Continue and observe.
- 👀 **`v = 0, ok = false`** — closed channel returns zero value

## Key Takeaway
**Channels synchronize goroutines through blocking.** Unbuffered channels require both sides to be ready. Buffered channels decouple sender and receiver until the buffer fills. Use the Goroutines panel to see what's blocked.
