class Solution {    
    public boolean isValid(String s) {
        if (s.length() % 2 != 0) return false;
        
        char lastChar = s.charAt(s.length() - 1);
        if (lastChar == '(' || lastChar == '[' || lastChar == '{') return false;
        
        Stack<Character> leftParentheses = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            char current = s.charAt(i);
            if (current == '(') {
                leftParentheses.push(')');
            } else if (current == '[') {
                leftParentheses.push(']');
            } else if (current == '{') {
                leftParentheses.push('}');
            } else if (leftParentheses.empty() || current != leftParentheses.pop()) {
                return false;
            }
        }
        
        if (leftParentheses.empty()) return true;
        return false;
    }
}