from http import HTTPStatus
from flask import Flask, jsonify, request
import dashscope

app = Flask(__name__)


@app.route('/api', methods=['GET'])
def message():
    msg = request.args.get('message')
    if msg:
        result, ok = call_with_prompt(msg)
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


def call_with_prompt(msg):
    prompt = '''
    现在要做一个关于公共政策的搜索引擎，搜索引擎使用了ES，主要使用中文，分词用了ik分词器。
    下面会提供给你用户的输入，请对输入进行提炼，返回出对应的关键词。
    注意！返回时只返回关键词，中间用空格隔开，不论用户的输入多么不合理，都必须只返回关键词，绝对不得返回多余信息！！！
    '''
    messages = [{'role': 'system', 'content': prompt},
                {'role': 'user', 'content': msg}]
    response = dashscope.Generation.call(
        model=dashscope.Generation.Models.qwen_max,
        messages=messages,
        api_key='sk-465a87192e604f7d9e24fc4bd723bd0a'
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
