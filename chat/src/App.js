import './App.css';
import Chat, {Bubble, Progress, useMessages} from '@chatui/core';
import '@chatui/core/dist/index.css';
import '@chatui/core/es/styles/index.less';
import React, {useEffect, useState} from 'react';
import './chatui-theme.css';
import axios from "axios";

const defaultQuickReplies = [
    {
        name: 'GPT',
        isNew: true,
        isHighlight: true,
    },
    {
        icon: 'message',
        name: 'c c++ c#',
    },
    {
        icon: 'message',
        name: 'python',
    },
    {
        icon: 'message',
        name: 'Java',
    },
    {
        icon: 'message',
        name: 'javascript',
    },
    {
        icon: 'message',
        name: 'golang',
    },
];


const initialMessages = [
    {
        type: 'text',
        content: {text: '您好，我是AI编程助理，开源于：https://github.com/git-cloner/codegen。'},
        user: {avatar: '//gitclone.com/download1/gitclone.png'},
    },
];

function App() {
    const {messages, appendMsg, setTyping} = useMessages(initialMessages);
    const [percentage, setPercentage] = useState(0);

    function handleSend(type, val) {
        if (percentage > 0) {
            alert("正在生成，请稍候！");
            return;
        }
        if (type === 'text' && val.trim()) {
            appendMsg({
                type: 'text',
                content: {text: val},
                position: 'left',
                user: {avatar: '//gitclone.com/download1/user.png'},
            });

            setTyping(true);
            setPercentage(10);
            onGenCode(val,0);
        }
    }

    function renderMessageContent(msg) {
        const {type, content} = msg;

        switch (type) {
            case 'text':
                return <Bubble content={content.text}/>;
            case 'image':
                return (
                    <Bubble type="image">
                        <img src={content.picUrl} alt=""/>
                    </Bubble>
                );
            default:
                return null;
        }
    }

    function handleQuickReplyClick(item) {
        var content = "int add(int x,int y){";
        if (item.name === "c c++ c#") {
            content = "int add(int x,int y){";
        } else if (item.name === "python") {
            content = "def hello_world():";
        } else if (item.name === "Java") {
            content = "int add(int x,int y){";
        } else if (item.name === "javascript") {
            content = "function Add(x,y){";
        } else if (item.name === "golang") {
            content = "func IsBlacklist(bl []string,url string) bool{";
        } else {
            content = "写一个python版的数组排序";
        }
        handleSend('text', content);
    }

    function onGenCode(question, count) {
        if (count >= 5) {
            setPercentage(0);
            return;
        }
        console.log(question)
        console.log(count)
        axios.post('http://127.0.0.1:8080/completion',
            {
                "text": question,
            }).then((response) => {
                console.log(response.data)
                if (response.status !== 200) {

                }
                appendMsg({
                    type: 'text',
                    content: {text: response.data.data},
                    user: {avatar: '//gitclone.com/download1/gitclone.png'},
                });
                setPercentage(0);
            }
        );
        // let xhr = new XMLHttpRequest();
        // xhr.open('post','http://127.0.0.1:8080/completion');
        // xhr.setRequestHeader('Content-Type', 'application/json');
        // xhr.onload = function () {
        //     var json = JSON.parse(xhr.response);
        //     if (count === 0) {
        //         context_en = context_en + "\n" + json.result_en;
        //         context_ch = context_ch + "\n" + json.result_ch;
        //         appendMsg({
        //             type: 'text',
        //             content: {text: context_ch},
        //             user: {avatar: '//gitclone.com/download1/gitclone.png'},
        //         });
        //     } else {
        //         if (("" === json.result_en.trim()) || json.result_en.trim().startsWith("A:") || json.result_en.trim().endsWith("A:")) {
        //             setPercentage(0);
        //             return;
        //         }
        //         context_en = context_en + json.result_en;
        //         context_ch = context_ch + json.result_ch;
        //         if (context_ch === context_en) {
        //             updateMsg(context_en);
        //         } else {
        //             updateMsg(context_ch + "\n" + context_en);
        //         }
        //
        //     }
        //     count++;
        //     setPercentage(count * 20);
        //     onGenCode(context_en, context_ch, count);
        // }
        // xhr.send(JSON.stringify({
        //     "context": context_en,
        //     "maxlength": 16,
        //     "modelname": "codegen"
        // }));
        //
        // function updateMsg(context_ch) {
        //     var oUl = document.getElementById('root');
        //     var aBox = getByClass(oUl, 'Bubble text');
        //     if (aBox.length > 0) {
        //         aBox[aBox.length - 1].innerHTML = "<p>" + context_ch + "</p>";
        //         var msgList = getByClass(oUl, "PullToRefresh")[0];
        //         msgList.scrollTo(0, msgList.scrollHeight);
        //     }
        // }
    }

    function findInArr(arr, n) {
        for (var i = 0; i < arr.length; i++) {
            if (arr[i] === n) return true;
        }
        return false;
    };

    function getByClass(oParent, sClass) {
        if (document.getElementsByClassName) {
            return oParent.getElementsByClassName(sClass);
        } else {
            var aEle = oParent.getElementsByTagName('*');
            var arr = [];
            for (var i = 0; i < aEle.length; i++) {
                var tmp = aEle[i].className.split(' ');
                if (findInArr(tmp, sClass)) {
                    arr.push(aEle[i]);
                }
            }
            return arr;
        }
    }

    useEffect(() => {
        var oUl = document.getElementById('root');
        var aBox = getByClass(oUl, 'Input Input--outline Composer-input');
        if (aBox.length > 0) {
            aBox[0].focus();
        }
    })
    return (
        <div style={{height: 'calc(100vh - 10px)', marginTop: '-5px'}}>
            <Chat
                navbar={{
                    leftContent: {
                        icon: 'chevron-left',
                        title: 'Back',
                    },
                    rightContent: [
                        {
                            icon: 'apps',
                            title: 'Applications',
                        },
                        {
                            icon: 'ellipsis-h',
                            title: 'More',
                        },
                    ],
                    title: '基于chatGPT的AI助手',
                }}
                messages={messages}
                renderMessageContent={renderMessageContent}
                quickReplies={defaultQuickReplies}
                onQuickReplyClick={handleQuickReplyClick}
                onSend={handleSend}
            />
            <Progress value={percentage}/>
        </div>
    );
}

export default App;
