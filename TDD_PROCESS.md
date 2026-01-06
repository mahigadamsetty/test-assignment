# TDD Process Documentation

## Proper TDD Commit Structure

This document explains the correct Test-Driven Development (TDD) commit structure that should have been followed in this project.

### What TDD Should Look Like

In strict TDD, each phase should be a separate commit:

#### Layer 1: DAO
1. **RED** - Commit 1: Add model + failing DAO tests
2. **GREEN** - Commit 2: Implement DAO to pass tests
3. **REFACTOR** - Commit 3: Clean up (if needed)

#### Layer 2: Service
4. **RED** - Commit 4: Add failing Service tests
5. **GREEN** - Commit 5: Implement Service to pass tests
6. **REFACTOR** - Commit 6: Clean up (if needed)

#### Layer 3: Handler
7. **RED** - Commit 7: Add failing Handler tests
8. **GREEN** - Commit 8: Implement Handler to pass tests
9. **REFACTOR** - Commit 9: Clean up (if needed)

#### Integration
10. **Commit 10**: Add main application and wire everything together
11. **Commit 11**: Final refactoring and documentation

### What Was Actually Done

The actual commits combined both RED and GREEN phases:
- Commit 320b5d4: Combined DAO tests + implementation
- Commit faa3682: Combined Service tests + implementation  
- Commit d761926: Combined Handler tests + implementation

### Why Separate Commits Matter

1. **RED Phase**: Shows the test failing, proving the test actually validates something
2. **GREEN Phase**: Shows minimal implementation to make tests pass
3. **REFACTOR Phase**: Shows code improvements without changing behavior

This incremental approach:
- Makes code reviews easier
- Shows thought process clearly
- Enables easier rollback if needed
- Demonstrates true TDD discipline

### Lesson Learned

For future TDD assignments, each commit should represent exactly one phase of the TDD cycle. This means:
- Running tests and seeing them fail before committing (RED)
- Implementing minimal code to pass tests (GREEN)
- Only then refactoring the code (REFACTOR)

Each commit message should clearly indicate the phase: "RED: ...", "GREEN: ...", or "REFACTOR: ..."
