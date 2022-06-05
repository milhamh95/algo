const generateParentheses = require('./generate_parentheses')



describe("success", function (){
    test("n = 3", function () {
        let res = generateParentheses(3)
        expect(res.length).toBe(5)
    })

    test("n = 1", function () {
        let res = generateParentheses(1)
        expect(res.length).toBe(1)
    })
})
