// @file-description This is an example of a JavaScript file.

// @project-title JavaScript Example
// @project-version 1.0.0
// @project-description A simple example of a JavaScript file working with godoc documentation.

// @description This is the main function.
function main() {
    console.log(add(1, 2) + " is the sum of 1 and 2.");
}

// @description This function adds two numbers together.
// @param a - The first number.
// @param b - The second number.
// @return The sum of the two numbers.
// @example add(1, 2)
// @complexity O(1)
// @author John Doe
function add(a, b) {
    return a + b;
}

main();