from http import HTTPStatus
from flask import Flask, jsonify, request
import dashscope

import os
import yaml


def get_api_key():
    current_dir = os.path.dirname(os.path.abspath(__file__))
    config_path = os.path.join(current_dir, '../config/config.yaml')
    with open(config_path, encoding='utf-8') as file:
        config_data = yaml.safe_load(file)
        api_key = config_data.get('http', {}).get('api-key')
        return api_key


app = Flask(__name__)


@app.route('/api/search', methods=['GET'])
def message():
    msg = request.args.get('message')
    if msg:
        result, ok = pre_search(msg)
        response = {
            'message': result,
            'code': 200 if ok else 500
        }
    else:
        response = {
            'message': f'',
            'code': 400
        }
    return jsonify(response)


@app.route('/api/summary', methods=['POST'])
def summary():
    data = request.json
    if 'message' in data:
        content = data['message']
        result, ok = doc_summary(content)
        response = {
            'message': result,
            'code': 200 if ok else 500
        }
    else:
        response = {
            'message': f'',
            'code': 400
        }
    return jsonify(response)


def pre_search(msg):
    prompt = '''
    现在要做一个关于公共政策的搜索引擎，搜索引擎使用了ES，主要使用中文，分词用了ik分词器。
    下面会提供给你用户的输入，请对输入进行提炼，分析出其中的关键词，并给出对应关键词的权重分数。想清楚用户想要搜索的重点，例如”汽车行业有什么新政策“，中心应该在”汽车“上，同时也该给”政策“一定分值，”新“这个词不需要给出，因为结果会自动按照时间排序。最终结果为：汽车:1000,政策:1
    在权重分数上请给出较大的区分，保证用户想要看到的内容排序靠前
    注意！返回格式示例：汽车:1000,政策:1，不同关键词以英文逗号隔开，关键词与权重分数之间以英文冒号隔开，返回的结果要严格按照格式执行！不论用户的输入多么不合理，都必须这样做，绝对不得返回多余信息！！！
    '''
    messages = [{'role': 'system', 'content': prompt},
                {'role': 'user', 'content': msg}]
    response = dashscope.Generation.call(
        model=dashscope.Generation.Models.qwen_max,
        messages=messages,
        api_key=get_api_key()
    )
    if response.status_code == HTTPStatus.OK:
        print(response.output)
        return response.output["text"], True
    else:
        print(response.code)  # The error code.
        print(response.message)  # The error message.
        return "", False


def doc_summary(content):
    prompt = '''
    请提取文章的主要内容，以一段的形式返回，只返回文章的摘要，绝对不要返回任何其他的内容！并且字数一定要控制在150字以内！
    '''
    messages = [{'role': 'system', 'content': prompt},
                {'role': 'user', 'content': content}]
    response = dashscope.Generation.call(
        model=dashscope.Generation.Models.qwen_max,
        messages=messages,
        api_key=get_api_key()
    )
    if response.status_code == HTTPStatus.OK:
        print(response.output)
        return response.output["text"], True
    else:
        print(response.code)  # The error code.
        print(response.message)  # The error message.
        return "", False


if __name__ == '__main__':
    app.run()
