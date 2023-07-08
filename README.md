# What's This For

This repo holds the code for [MSDS-431 Learning Data Engineering with Go](https://msdsgo.netlify.app/data-engineering-with-go/) Week 4 discussion.

# Discussion Prompt

### Week 4 Discussion. Pros and Cons of Generics

Generics were introduced into the Go language for the first time in Go version 1.18Links to an external site..

Read about generics in the revised version of chapter 15 of Jon Bodner's (2021) Learning Go: An Idiomatic Approach to Real-World Go Programming, available from the book's websiteLinks to an external site. and attached here for your convenience: chapter15_learningGo.pdfDownload chapter15_learningGo.pdf

Go version 1.0 was released more than ten years ago and Go programmers have gotten along fine without the generics all that time.

For a short list of generics advantages and disadvantages, see:  Generics can make your Go code slower.Links to an external site.

Post a short Go program that uses generics, as well as a short program that accomplishes the same task without generics. How easy was it to write each program? 

Include CPU timing code with your programs. Which program runs faster?  

From your reading and coding experiment, describe advantages and disadvantages of generics.

How do you feel about using generics with Go?

# My Response

I wrote a quick program that takes a slice of numbers as an input and returns the maximum from the slice.  Here's the code:

\<see the main.go file here on github.  For the discussion board I actually posted it>

As you can see, this necessitated the creation of three functions: one to return the maximum from a slice of ints, one for the maximum from a slice of float64s, and one from a generic list of types that could be Ordered.

In reading through this, I found it interesting that even the concept of generics isn't universally generic: there are still these constraint types that require forethought for exactly what sort of inputs can be given.  I guess that's not surprising given the thesis of go being to strongly type and ensure a relatively locked-down code environment.

For what it's worth, the generic function ran 7-10x slower than the typed functions (60-70µs generic vs 7-10µs typed).  This is a massive performance hit.

Which then raises the question of the pros and cons of using generics.  It's obvious: do you want your code to be relatively more streamlined or do you want your execution time to be faster?  The burden of figuring out your data either sits with the coder or with the compiler, so which one do you trust?  If we use generics then methods can be reusable which reduces the chance for coding errors.  If we don't use generics then we have to make sure our code is airtight but it makes life a lot easier at runtime.