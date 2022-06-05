var generateParenthesis = function(n) {
    let stack = []

    generate(stack, "", 0, 0, n)

    return stack
};

function generate(stack, s, open, close, n) {
    if (open == n && close == n) {
        stack.push(s)
        return
    }

    if (open < n) {
        generate(stack, s + "(", open + 1, close, n)
    }

    if (close < open) {
        generate(stack, s + ")", open, close + 1, n)
    }
}

module.exports = generateParenthesis