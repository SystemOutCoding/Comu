from __future__ import print_function
import numpy as np
import tensorflow as tf

#5,42
tf.set_random_seed(182)
np.random.seed(1)

#input data
xy = np.loadtxt("zebal_final_norm.csv", delimiter=",", dtype=np.float32)
#train/test split parameter
datasplit=560

#train data
train_x = np.array(xy[:datasplit,1:])
train_y = np.array(xy[:datasplit,0])

#test data
test_x = np.array(xy[datasplit:,1:])
test_y = np.array(xy[datasplit:,0])

# Parameters
learning_rate = 0.03
display_step = 1

# Network Parameters
n_hidden_1 = 5 # 1st layer number of neurons
n_hidden_2 = 5 # 2nd layer number of neurons
n_hidden_3 = 5 # 3rd layer number of neurons
n_hidden_4 = 5 # 3rd layer number of neurons
n_hidden_5 = 5 # 3rd layer number of neurons
n_input = len(np.transpose(train_x)) # data input
n_output = 1 # total output


# Store layers weight & bias
weights = {
    'h1': tf.Variable(tf.random_normal([n_input, n_hidden_1])),
    'h2': tf.Variable(tf.random_normal([n_hidden_1, n_hidden_2])),
    'h3': tf.Variable(tf.random_normal([n_hidden_2, n_hidden_3])),
    'h4': tf.Variable(tf.random_normal([n_hidden_3, n_hidden_4])),
    'h5': tf.Variable(tf.random_normal([n_hidden_4, n_hidden_5])),
    'out': tf.Variable(tf.random_normal([n_hidden_5, n_output]))
}
biases = {
    'b1': tf.Variable(tf.random_normal([n_hidden_1],0.5,1.0)),
    'b2': tf.Variable(tf.random_normal([n_hidden_2],0.5,1.0)),
    'b3': tf.Variable(tf.random_normal([n_hidden_3],0.5,1.0)),
    'b4': tf.Variable(tf.random_normal([n_hidden_4],0.5,1.0)),
    'b5': tf.Variable(tf.random_normal([n_hidden_5],0.5,1.0)),
    'out': tf.Variable(tf.random_normal([n_output]))
}

# Create model
def multilayer_perceptron(x):

    layer_1 = tf.add(tf.matmul(x, weights['h1']), biases['b1'])
    #2,3,4,5 layer에는 relu적용
    layer_2 = tf.add(tf.matmul(layer_1, weights['h2']), biases['b2'])
    layer_3 = tf.nn.relu(tf.add(tf.matmul(layer_2, weights['h3']), biases['b3']))
    layer_4 = tf.nn.relu(tf.add(tf.matmul(layer_3, weights['h4']), biases['b4']))
    layer_5 = tf.nn.relu(tf.add(tf.matmul(layer_4, weights['h5']), biases['b5']))
    out_layer = tf.matmul(layer_5, weights['out']) + biases['out']
    return out_layer

# Construct model
hypothesis = multilayer_perceptron(train_x)
test_hypothesis = multilayer_perceptron(test_x)

# Define loss
loss_op = tf.reduce_mean(tf.square(hypothesis - train_y))
test_loss=tf.reduce_mean(tf.square(test_hypothesis-test_y))

#define optimizer
optimizer = tf.train.AdamOptimizer(learning_rate=learning_rate)
train_op = optimizer.minimize(loss_op)

# befor starting, initialize variables
init = tf.initialize_all_variables()

#launch
sess = tf.Session()
sess.run(init)


print('learning rate:{}'.format(learning_rate))
print('input to output : {},{},{},{},{}'.format(n_input,n_hidden_1,n_hidden_2,n_hidden_3,n_output))

# fit the line
for step in range(20001):
    sess.run(train_op)
    if step % 200 == 0:
        print('-------------------------')
        print('{} steps'.format(step))
        print('    train:  {}'.format(sess.run(loss_op)))
        print('    test:   {}'.format(sess.run(test_loss)))

"""
#경기차 출력 
for step in range(len(train_x)):
    print('{},{},'.format(step+1,round(sess.run(162*0.4506*tf.subtract(test_hypothesis[step],test_y[step]))[0]+0.2654,0)))
    print('{},{},'.format(step + 1,round(sess.run(162 * 0.4506 * tf.subtract(test_hypothesis[step], test_y[step]))[0],0)))
"""