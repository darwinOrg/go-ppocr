#!/usr/local/bin/python3
# -*- coding: UTF-8 -*-

import json
import sys

from paddleocr import PaddleOCR

img_path=sys.argv[1]
out_path=sys.argv[2]

# Paddleocr目前支持的多语言语种可以通过修改lang参数进行切换
# 例如`ch`, `en`, `fr`, `german`, `korean`, `japan`
ocr = PaddleOCR(use_angle_cls=True, lang="ch", show_log=False)
result = ocr.ocr(img_path, cls=False)
result_json = json.dumps(result, ensure_ascii=False)
print(result_json)
with open(out_path, 'w') as fo:
    fo.write(result_json)
