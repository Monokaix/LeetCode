/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    //递归先序遍历
    public String serialize(TreeNode root) {
        if (root == null) {
			return "#!";
		}
		String res = root.val + "!";
		res += serialize(root.left);
		res += serialize(root.right);
		return res;
        
    }

    //先于处理数据放入队列再递归处理
    public TreeNode deserialize(String data) {
        String[] values = data.split("!");
		Queue<String> queue = new LinkedList<String>();
		for (int i = 0; i != values.length; i++) {
			queue.offer(values[i]);
		}
		return reconPreOrder(queue);
    }
    public static TreeNode reconPreOrder(Queue<String> queue) {
		String value = queue.poll();
		if (value.equals("#")) {
			return null;
		}
		TreeNode head = new TreeNode(Integer.valueOf(value));
		head.left = reconPreOrder(queue);
		head.right = reconPreOrder(queue);
		return head;
	}
}