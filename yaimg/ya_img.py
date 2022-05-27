import re

import requests
from bs4 import BeautifulSoup
from urllib.parse import quote


def try_parse_img_size(string: str, img_size_split_char='×') -> (int, int):
    sizes = re.findall(rf"[0-9]+{img_size_split_char}[0-9]+", string)
    return tuple(map(int, sizes[0].split(img_size_split_char))) if len(sizes) == 1 else (0, 0)


def get_pic_src_with_sizes(pic_url: str, img_size_split_char='×') -> (list, list, (int, int)):
    """ :returns list[src_url: str, (width: int, height: int)],
                 list[warnings: str],
                 (original_image_width: int, original_image_height: int)"""

    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36'}

    resp = requests.get(f"https://yandex.ru/images/search?rpt=imageview&url={quote(pic_url)}", headers=headers)
    if resp.status_code != 200:
        return list(), [resp.text], (0, 0)
    page_soup = BeautifulSoup(resp.text, 'html.parser')

    report = list()
    warnings = list()
    for button_to_other_source in page_soup.find_all('a', class_="Button2"):
        spans = button_to_other_source.find_all('span')
        if len(spans) == 1 and img_size_split_char in spans[0].string:
            try:
                src_url = button_to_other_source['href']
                report.append((src_url, try_parse_img_size(spans[0].string)))
            except KeyError as err:
                warnings.append(str(err))

    original_image_size = (0, 0)
    original_image_size_infos = page_soup.find_all('div', 'CbirPreview-ImageSize')
    if len(original_image_size_infos) != 1:
        warnings.append(f"len(original_image_size_infos) != 1, {str(original_image_size_infos)}")
    else:
        original_image_size = try_parse_img_size(original_image_size_infos[0].text.split(':')[-1])

    return report, warnings, original_image_size


if __name__ == "__main__":
    from argparse import ArgumentParser
    from sys import argv
    import json

    parser = ArgumentParser(description='A script to get sources of an image via Yandex Reverse Image Search, '
                                        'returns json format:   {'
                                        '                         "src": [[url_str_1, [height_int_1, width_int_1]],..],'
                                        '                         "warn": [warn_str_1,..], '
                                        '                         "sizes": [height_int, width_int]'
                                        '                       }')
    parser.add_argument('-u', '--url', metavar='URL', type=str,
                        help='URL to original picture')
    parser.add_argument('-c', '--split-char', default='×', required=False, type=str,
                        help='(optional) a character that Yandex uses to split image sizes')

    namespace = parser.parse_args(argv[1:])

    src, warnings, original_image_size = get_pic_src_with_sizes(namespace.url, namespace.split_char)
    formatted_src = list(map(lambda x: { "url": x[0], "width": x[1][0], "height": x[1][1]}, src))
    print(json.dumps({'src': formatted_src, 'warn': warnings, 'sizes': original_image_size}))
