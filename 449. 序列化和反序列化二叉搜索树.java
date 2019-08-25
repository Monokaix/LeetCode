
//二叉搜索树插入
void BST_insert(TreeNode *node, TreeNode *insert_node){
	if (insert_node->val < node->val){
		if (node->left){
			BST_insert(node->left, insert_node);
		}
		else{
			node->left = insert_node;
		}
	}
	else{
		if (node->right){
			BST_insert(node->right, insert_node);
		}
		else{
			node->right = insert_node;
		}
	}
}

void change_int_to_string(int val, std::string &str_val){
	std::string tmp;
	while(val){
		tmp += val % 10 + '0';
		val = val / 10;
	}
	for (int i = tmp.length() - 1; i >= 0; i--){
		str_val += tmp[i];
	}
	str_val += '#';
}
//先序遍历
void BST_preorder(TreeNode *node, std::string &data){
	if (!node){
		return;
	}
	std::string str_val;
	change_int_to_string(node->val, str_val);
	data += str_val;
	BST_preorder(node->left, data);
	BST_preorder(node->right, data);
}

class Codec {
public:
    std::string serialize(TreeNode* root) {
    	std::string data;
        BST_preorder(root, data);
        return data;
    }
    TreeNode *deserialize(std::string data) {
    	if (data.length() == 0){
	    	return NULL;
	    }
    	std::vector<TreeNode *> node_vec;
    	int val = 0;
    	for (int i = 0; i < data.length(); i++){
	    	if (data[i] == '#'){
	    		node_vec.push_back(new TreeNode(val));
	    		val = 0;
	    	}
	    	else{
	    		val = val * 10 + data[i] - '0';
	    	}
	    }
	    for (int i = 1; i < node_vec.size(); i++){
    		BST_insert(node_vec[0], node_vec[i]);
    	}
    	return node_vec[0];
    }
};
