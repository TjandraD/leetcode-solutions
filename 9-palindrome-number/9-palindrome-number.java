class Solution {
    public boolean isPalindrome(int x) {
        String input = Integer.toString(x);
            String[] arrayInput = input.split("");

            for (int i = 0; i < arrayInput.length; i++) {
                int position = arrayInput.length - i - 1;
                if (!Objects.equals(arrayInput[i], arrayInput[position])) return false;
            }

            return true;
    }
}