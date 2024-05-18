// @description This is an example of a JavaScript file.

// @title JavaScript Example
// @version
// @description A simple example of a JavaScript file working with godoc documentation.
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