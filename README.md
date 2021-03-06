# Cryptica


## Context

[Cryptica](https://play.google.com/store/apps/details?id=pixibots.games.cryptica.com) is a mobile game developed by Pixibots. I have no affiliation with the company other than being a satisfied customer. Solving a puzzle in Cryptica involves moving tiles on a 7x5 board to match a desired configuration. All tiles move symultaneously; their movements are, however, constrained by the board edges, and possibly also by some additional (unmovable) barriers.

## Description

This Go package can search efficiently for solutions of Cryptica and Cryptica-like puzzles, in either depth-first or breadth-first fashion. Moreover, it can work on boards of arbitrary size (with some caveats).

## History

I wrote this software in early 2017 as part of a hands-on effort to learn Go and its development tools. I would like to thank Wojciech Smalinski at Google for introducing me to this weird but powerful programming language.

## Documentation

https://godoc.org/github.com/rvirga/cryptica

## Note

If you came here just looking for the solutions of the game, these are listed below. They have been generated
automatically by this software in under an hour, on some pretty dated hardware.

* **AJAW**
  1. ↑ ↑ ↑ ↑ → → →
  2. ↑ ↑ → → → ↓ ↓ ↓ ↓ →
  3. ← ← ↑ ← ← ←
  4. ↓ ↓ ← ← ↑ ↑ ↑ ↑
  5. → → → → → ↓ ← ← ← ← ← ←
  6. ↑ → → → → ↓ ↓ ← ← ← ← ← ↑
  7. ↑ ↑ → → → ↑ ↓ ↓ ↓ ← ← ← ← ← ← ↓ → → → → →
  8. ↑ → → ↑ ← ← ← ↓ ↓ ↓ ↓
  9. ↑ ↑ ← ← ↓ ↓ ↓ ↓ → → → → → → ← ← ←
  10. ↓ ↓ ↓ ← ← ← ↓ → → → → ↑ ↑ ↑ ↑
  11. ↑ ← ← ↑ ↑ ↓ → → → →
  12. ↑ → → → → → ← ← ↑ ↑ ↓ ←
  13. ← ← ← ↓ ← ↓ ← ← ↑ → → ↑ ← ↓ ↓ ↓
  14. → → → → → → ↑ ↑ ← ↓ → ← ↓ → ← ↓ → ← ↓ ← ←
  15. ↑ ↑ ↓ ← ← → ↑ ↑ ← ↓ ↓ ↓ ↓ ← → →
  16. ← ↓ → → ↑ ← ← ↓ ↓ ↓
  17. ↓ ← ← ← ← ← ← ↑ ↑ ↑ → → → → → → ↑ ← ← ← ← ←
  18. ↑ ← ← ↑ ← ↓ ↓ ↓ ← ← ← →
  19. → → → ← ↑ → ↓ ↓ ↓ ↓ ← ←
  20. ↑ ← ↓ ↓ ↑ → → → ↓ ←
  21. ↑ ← ← ↑ → → → → ↓ ↓ ↓
  22. ↓ → → → → → ↑ ↑ ← ←
  23. ← ← ↓ → → → → → → ← ↓ ↓ ↓
  24. → ↑ ↑ → → ↑ ↑ ← ↓ ↓ ← ← ← ← ↓ ↓ → → → → → ↑
  25. ↓ ← ← ← ← → → ↑ ↑
  26. ↓ ↓ ← ↓ ← ↑ ↑ ← ← ← ↓ ↓ → ↑
  27. → → → → ↓ ↓ ↓ ↑ ↑ ← ↓ ↓ → ↓ ← ←
  28. ↓ ↓ ← ← ↓ → → → → ↑ ↑ ↑ ↑
  29. ↓ ↓ ↓ → → → ↓
  30. → → ↑ → ↑ → ↑ ← ↓ ↓ ↓ ← ↑ ←
* **B'ALAM**
  1. ↓ ← ← ← ← ↑ ↑ ↑ ↑ ↓ ← → → ↓ ← ↑ →
  2. ← ↑ → → ↓ ← ↓ ↓ ← ← ↑ → → → → →
  3. ↑ → ↓ ↓ ↓ ← ← ← ←
  4. ↑ ← ↑ → → → → ↓ ↓ ↓ ↓ ←
  5. ↑ ← ↓ → → → → → ↑ ↓ → ↑ ← ← ← ← ← → →
  6. ↑ ↑ ↑ ← ← ← ↑ ↓ ↓ → → → ↑ ↑ ↓ ← ← ← ↑ → → → ↓ ↓
  7. ↓ ↓ → ↓ ↓ ↑ ↑ → ↑ ↑ ↓ ↓ ← ← ← ← ← ↓ ↑ →
  8. → → → ↓ ↓ ↓ ↓ ↑ ↑ → ↑ ↑ ↓ ↓ ← ← ← ← ← ↓ ↓ ↑ ↑
  9. ↑ ↑ ← ↑ ↓ ← ← ↑ → ↑ ↑ ← → ↓ ↓ ← ↑ ↑ → → → → → → ↓ ← ←
  10. ↑ → → → ↑ ↑ ↓ ← ← ← ← ← ← ↑ ↑ ↓ ← → → →
  11. → ↓ ↓ ↓ ← ← ↓ → ↑ ← ↑ ↑ ← ↓ ←
  12. ↓ ↓ ↓ ↓ ← ↑ ↑ ↑ ↑ → → → → →
  13. ↑ ↑ → → ↑ → ↓ ↓ ← ↓ → ↑ ← ← ← ←
  14. ↑ → → ↓ ↓ → → ↑ → ← ← ↑
  15. ↑ ← ← ↑ ← ↓ → → ↑ → ↓ ↓ ↓
  16. ← ↓ → → → ↓ → ↑ ↓ ↓ ← ← ← ← ← ↑ ↑
  17. → ↑ ← ← ↑ → ↓ ↓ ↓ ↓ ↑ ↑ ←
  18. → → ↓ ↓ ← ↓ → ↓ ↓ ↓ ← ← ← ↓ ↓ → ↑ ↑ ←
  19. ↑ → → → → ↑ ↑ ↓ → ↓ → ↑ ↑ ↓ → → →
  20. ← ↓ ← ↓ → → → ↑ ↓ ↓ ←
  21. ↑ → ↓ → → ↓ → ← ← ↑ ↑ ← ←
  22. → → → → → ↓ ← ← ← ↓ ← ↓ ← ↓ ← → ↓ ← ↓ → ↓ → ↓ → ↑ → → ↓
  23. → → ↓ → ↓ ← ← ← ← ↓ ←
  24. ↑ ← ← → ↓ ↓ ↓ ↓ ↑ ↑ ↑ ← ↑ ↓ ↓ ↓ → → → ↑ → ↑ ↑ ← ← ← ↓ ← ↓
  25. → → ↑ → ↓ ↓ ↓ → → → → ← ← ↑ ← ↓ → →
  26. → ↑ ↓ → → ↑ ↑ ↓ → ← ↑ ← ← ← ← ←
  27. ↑ ← ↑ ↓ ↓ → → ↑ ↑ ↑ ← ← ←
  28. ← → ↑ → ← ↑ → → ↓
  29. ↑ ← ↓ ↑ ← ↑ ↑ ↑ → ↓ ↓ ↓ ↓ → ←
  30. ↑ → ↓ → ↓ → ↓ ← ← ← ← ↑ ←
* **K'UK**
  1. → ↑ ↑ ← ↑ ↑ → → ↓ ← ← ← ↓ ↓ ↓ → →
  2. → ↓ → ↑ → → ↑ ↑ ← ↑ → → ↓ → ↑ ← ← ← ←
  3. ↓ → ↑ ↓ → ↑ ← ← ↑ → ↓ ← ← ↓ ← ↓ ← ↓ ← → ↓ ← ↓ → → → ↓ → →
  4. ← ← ↓ ← ← → → → ↓ → ↑ ↑ ↑ ← → ↑ ←
  5. ↓ ↓ ← ← ← ↑ ↑ ↑ ↑ ↓ ↓ → ↓ ↓ ↑ ↑ → → ↓ ↓
  6. → ↓ ← → ↓ → ↓ ↓ ← → ↓ → ↓ ↓ ↑ ← ↑ ↑ ↑ ↓
  7. ↓ ↑ → → ↓ ↓ ← ← ↓ ↑ ↑ ← ↓
  8. ↓ ← ↓ → ↓ → ↑ → ↓ → ↓ ← ← ← ↓ → ↑ ↑ ← ← ↑ → ↑ → →
  9. ↑ ← ↑ ← ← ↓ ↓ → → ←
  10. ↑ ← ← ↑ ↑ ← ↓ ↓ ↓ → → ↓ ← ← ↑ ↑ ↓ → → →
  11. ← ↓ → → → ↓ ← ← ← ← ← ↓ → ↑ ↑ ↑ ↑ ← → → ←
  12. ← ↑ ↓ ← ↓ ← ↑ ↓ ← ↑ ↑ ↑ → ↑ ↓ → → → → → ↓ ↓ ↓
  13. ↑ → ↓ ← ↑ ← ← ↑ → → → → ↓ ↓ ← ← ← ← ← ↑
  14. ↑ ↑ → → → → ↓ ↓ ↓ ← ← ← → ↑ ←
  15. ↓ ↓ → ↓ ↑ ← ← ↓ ↑ ↑ ↓ ← ↓
  16. ↑ → ← ↑ ↑ → → → → ↓ ↓ ↓ ↓ ← ← ← ← ←
  17. → ↑ ↑ ← ↓ ↓ ← ↓ → → ↑ ← → ↓ → ← ↓ → →
  18. ↑ ← ↓ → → → → → → ← ↓ → ← ↓ → ↑ ↑ ↑ ↑ ←
  19. ↑ ↑ → → → → ← ↑ → ← ← ↑ → → ↓ ↓ ↓ ↑ ↑
  20. ← ↑ ↑ → → → ↓ ← ← ↓ ← ← ← ← ↑ → ↑ ← ↑ → ↑ →
  21. ← ↓ ← ← ← ↑ → → ↑ ↑ ↑ ↓ ← ← ↑ ↓ ← → ↑ → ↑
  22. ↓ ↓ ← ← ↓ ← ↑ ↑ ← ↑ → → → →
  23. ↓ ↓ ← ← ↑ ← ↑ → ↑ ↑ ↓ → → → ↓ → ↑ ↑
  24. ↓ → ↑ ↓ → → ↑ ← ↑ → → ↑ ← ↓ → ↓ ← ← ← ↓ ↓ ↓
  25. ↑ → ↓ ↓ → ↑ ← ← ↑ ↓ ← ← ↑ ↓ ← ↑ ↑
  26. ← ↑ ← ← ↓ ↓ ↓ ↑ → → → → → → ↓
  27. ← ↑ ↑ ← ← ↓ ↓ ← ↓ ↓ → ↓ → ↑ ↑ → ↓ ↓
  28. ↓ → ↑ → → ↓ ← ← ← ↑ ↓ ← ← ← ↑ ↑ ↑
  29. ← ↓ → ↑ ↑ → → ↓ ↓ ↓ ↓ ↑ ↑ ↑ ↓ → ↑ ↑ ↓ ← ← ↓ ↓ ↓ ← ← →
  30. → → → ↓ ↓ → ↓ ← ↓ → ↑ ← ← ← ← ← ↓ ↓ ← ↓ → ↑ → → →
* **CHAN**
  1. ← ↓ ← ↑ ← ↑ → → → → ↑ → → ↓ → ↑ ← ↑ → ↓
  2. ↓ ← → ↓ ↑ ← ← ↑ ↑ ↑ → ↓ → → → → → ↓ ↓ ← ← ← ← ← → ↓ ↑ ← ←
  3. ↑ ↑ ↑ ← ↓ ↓ → ← ↓ → → ↓ ← ← ↑ ← ↓ ↓ → → ↑ ↑ → ↓ ↓ → ←
  4. ↓ ↓ ← ↑ ← ← ↓ ↑ → ↓ ↓ ↓ ← ← →
  5. → → ↑ → ↓ ↓ ← ← ← ← → ↑ ↑
  6. ↓ ↓ → ↓ → ← ← ← ← → ↑ ↑ → → ← ← ↑ ↑ ← →
  7. ↓ ↑ ← ← ↓ ← ↑ → → → → ↓ → → ↓ ← ← ← ← ← ← ↓ ← ↑ ↑ ↑ ↑ ← ← ←
  8. ← ↑ ← → ↑ → → ← ↑ ← ↑ ↓ → → → ← ↑ → ↓ ↓
  9. ↑ ← → → ↓ ← ↓ ↓ ↓ ← ↓ → → ↓ → → ↑ ↑ ↑ ↑ ← → → ←
  10. ← ↓ → ↓ ← ← → ↓ ↑ → ← ↑ → → ↑
  11. ↓ ↓ ↑ ↑ ← ← ← ← ← ↓ ↓ ↑ → → → → → → ↓ ←
  12. ↓ → → → → ↓ → ↓ ← ← ↓ ← ← ↓ ← ↑ → ↓ ← ↑ ↑ →
  13. ← → → ↓ → → ↓ ↑ ← ← ← ↓ ← ← ← → ↓ ↑ ←
  14. ↓ ← ← ← ↑ → ↑ ← → ↑ → ↑ ↓ ← ← ↑ ↓ ← ← ↓ ↑ ← ↓ ↓ ↓ →
  15. ← ← ↓ → → ↑ ↑ → ↑ ← ← ↓ → ↑ → ← ↓
  16. ← ↑ ↓ ← ↑ ↑ ↓ → → → ↑ ← ← ← ↓ ← ↑ → → → → → ↓
  17. ↑ ← ↑ ↓ ← ↑ ↓ ← ← ↑ → ↓ → ↑ ↑ ← ← ↑ ↑ ↑
  18. ↑ ← ← ↓ ↓ → ↓ ← ↑ ↑ ← ↑ → ↑ ← ↓ ↓ ↑ → ↑ ↑ ↑ ← ← ↓
  19. ← ↑ ← ↓ → → ↑ ← ↓ ← ←
  20. ↑ → ↓ → ↑ ↑ ← ← ← ↑ ↑ ↑ ↓ ↓
  21. ↑ → ↓ ↓ ← → ↑ → ↑ ↑ ← ← ← ← ↓ →
  22. ← ← ↓ ↓ → ← ↑ → → → ↑ → ↓ → ↓ ↓ ↓ ←
  23. ↑ ← ↓ ← ↓ → → → → ↑ ← ← ↑ ↓ ← ↓
  24. ↑ ← ← ↓ → ↑ → → ↓ → → → ↓ ↓ ← ← ↓ ← ← ← ← → ↑ ↑ ↑ ↑
  25. ↑ ↑ → ↑ ← → → ↓ ← ← ← ↓ ← ← ← ← ↑ ↑
  26. ↓ ↓ ↓ ↑ → ↑ ↓ → → ← ↓ ↑ ← ← → ↑ ← ← ↓ ←
  27. ↓ ↓ ↓ ← ← ← ↑ ↑ ↑ ↑ → → → ↓ ← ↑ ← ← ↓ ↓
  28. → ← ↓ ← ↑ → ← ↓ → →
  29. ↑ ↑ ↓ → → ↓ ← ← ↑ ↑ → → → ↓
  30. ← ↓ → → → ← ↑ ↑ ← ← ← ← ↓ → → ↓ ↑ ↑ → → ↑ ← ← ← ↑ → → → ↓ ↓ ↓ ↓
