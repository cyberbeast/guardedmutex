# guardedmutex

Demonstration of a pattern that ensures that a guarded value can't be accessed without first acquiring a lock. Inspired by Patricio Whittinglow's [YT video](https://youtu.be/318EQBvB1No?si=v4ec3sHvrNo3FkDW) reviewing an article comparing Rust and Go.

While I've published this implementation as a perfectly usable/importable MIT licensed package, I hope you'll apply this Go proverb if you find yourself reaching for this functionality - "[a little copying is better than a little dependency](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s)".

I've extended the pattern in my library to explore underlying variable mutation when using a pointer type as a type constraint (still guarded/protected by the Mutex) as well as effective error handling when using this pattern.
