// 应用状态
const state = {
    jsonInput: '',
    language: 'golang',
    tags: ['json', 'mapstructure'],
    structName: 'User',
    packageName: 'models',
    addComments: true,
    usePointer: false,
    timeAsString: false,
    isValidJson: false,
    converting: false,
    outputCode: ''
};

// 可用标签
const availableTags = [
    { value: 'json', label: 'JSON' },
    { value: 'mapstructure', label: 'MapStructure' },
    { value: 'gorm', label: 'GORM' },
    { value: 'bson', label: 'BSON' },
    { value: 'yaml', label: 'YAML' },
    { value: 'xml', label: 'XML' },
    { value: 'validate', label: 'Validate' },
    { value: 'form', label: 'Form' }
];

// 示例数据
const examples = [
    {
        name: '用户信息',
        description: '简单的用户对象示例',
        json: JSON.stringify({
            "id": 1,
            "name": "John Doe",
            "email": "john@example.com",
            "age": 30,
            "active": true,
            "createdAt": "2023-12-20T10:30:00Z"
        }, null, 2)
    },
    {
        name: '产品信息',
        description: '包含嵌套对象的产品信息',
        json: JSON.stringify({
            "productId": "P001",
            "name": "Laptop",
            "price": 999.99,
            "inStock": true,
            "specs": {
                "cpu": "Intel i7",
                "ram": "16GB",
                "storage": "512GB SSD"
            },
            "tags": ["electronics", "computer", "laptop"]
        }, null, 2)
    },
    {
        name: '订单数据',
        description: '包含数组和嵌套对象的订单',
        json: JSON.stringify({
            "orderId": "ORD123",
            "customer": {
                "name": "Alice Smith",
                "email": "alice@example.com"
            },
            "items": [
                {
                    "productId": "P001",
                    "quantity": 2,
                    "price": 99.99
                },
                {
                    "productId": "P002",
                    "quantity": 1,
                    "price": 199.99
                }
            ],
            "total": 399.97,
            "status": "shipped"
        }, null, 2)
    }
];

// DOM元素引用
const elements = {
    jsonInput: null,
    jsonStatus: null,
    language: null,
    tagsContainer: null,
    tagSection: null,
    structName: null,
    packageName: null,
    packageGroup: null,
    addComments: null,
    usePointer: null,
    timeAsString: null,
    convertBtn: null,
    outputCode: null,
    copyBtn: null,
    downloadBtn: null,
    clearBtn: null,
    examplesContainer: null
};

// 初始化应用
function initApp() {
    // 获取DOM元素
    elements.jsonInput = document.getElementById('json-input');
    elements.jsonStatus = document.getElementById('json-status');
    elements.language = document.getElementById('language');
    elements.tagsContainer = document.getElementById('tags-container');
    elements.tagSection = document.getElementById('tag-section');
    elements.structName = document.getElementById('struct-name');
    elements.packageName = document.getElementById('package-name');
    elements.packageGroup = document.getElementById('package-group');
    elements.addComments = document.getElementById('add-comments');
    elements.usePointer = document.getElementById('use-pointer');
    elements.timeAsString = document.getElementById('time-as-string');
    elements.convertBtn = document.getElementById('convert-btn');
    elements.outputCode = document.getElementById('output-code');
    elements.copyBtn = document.getElementById('copy-btn');
    elements.downloadBtn = document.getElementById('download-btn');
    elements.clearBtn = document.getElementById('clear-btn');
    elements.examplesContainer = document.getElementById('examples-container');

    // 绑定事件
    bindEvents();

    // 初始化UI
    renderTags();
    renderExamples();
    updateUI();

    // 加载第一个示例
    if (examples.length > 0) {
        loadExample(examples[0]);
    }
}

// 绑定事件
function bindEvents() {
    // JSON输入监听
    elements.jsonInput.addEventListener('input', (e) => {
        state.jsonInput = e.target.value;
        validateJson();
        updateUI();
    });

    // 语言选择监听
    elements.language.addEventListener('change', (e) => {
        state.language = e.target.value;
        updateTagSection();
        updatePackageGroup();
        updateUI();
    });

    // 配置选项监听
    elements.structName.addEventListener('input', (e) => {
        state.structName = e.target.value;
        updateUI();
    });

    elements.packageName.addEventListener('input', (e) => {
        state.packageName = e.target.value;
        updateUI();
    });

    elements.addComments.addEventListener('change', (e) => {
        state.addComments = e.target.checked;
        updateUI();
    });

    elements.usePointer.addEventListener('change', (e) => {
        state.usePointer = e.target.checked;
        updateUI();
    });

    elements.timeAsString.addEventListener('change', (e) => {
        state.timeAsString = e.target.checked;
        updateUI();
    });

    // 按钮点击监听
    elements.convertBtn.addEventListener('click', convertJson);
    elements.copyBtn.addEventListener('click', copyToClipboard);
    elements.downloadBtn.addEventListener('click', downloadCode);
    elements.clearBtn.addEventListener('click', clearOutput);
}

// 渲染标签选择器
function renderTags() {
    elements.tagsContainer.innerHTML = '';

    availableTags.forEach(tag => {
        const isActive = state.tags.includes(tag.value);
        const tagEl = document.createElement('span');
        tagEl.className = `tag ${isActive ? 'active' : ''}`;
        tagEl.textContent = tag.label;
        tagEl.dataset.value = tag.value;
        tagEl.addEventListener('click', () => toggleTag(tag.value));
        elements.tagsContainer.appendChild(tagEl);
    });
}

// 渲染示例
function renderExamples() {
    elements.examplesContainer.innerHTML = '';

    examples.forEach((example, index) => {
        const exampleEl = document.createElement('div');
        exampleEl.className = 'example-card';
        exampleEl.innerHTML = `
            <h3>${example.name}</h3>
            <p>${example.description}</p>
        `;
        exampleEl.addEventListener('click', () => loadExample(example));
        elements.examplesContainer.appendChild(exampleEl);
    });
}

// 验证JSON
function validateJson() {
    if (!state.jsonInput.trim()) {
        state.isValidJson = false;
        return;
    }

    try {
        JSON.parse(state.jsonInput);
        state.isValidJson = true;
    } catch (error) {
        state.isValidJson = false;
    }
}

// 更新UI
function updateUI() {
    // 更新JSON状态
    updateJsonStatus();

    // 更新按钮状态
    updateButtons();

    // 更新状态
    updateStateFromInputs();
}

// 更新JSON状态显示
function updateJsonStatus() {
    if (!state.jsonInput.trim()) {
        elements.jsonStatus.textContent = '等待输入...';
        elements.jsonStatus.className = 'status status-waiting';
        return;
    }

    if (state.isValidJson) {
        elements.jsonStatus.textContent = '✓ JSON有效';
        elements.jsonStatus.className = 'status status-valid';
    } else {
        elements.jsonStatus.textContent = '✗ JSON无效';
        elements.jsonStatus.className = 'status status-invalid';
    }
}

// 更新按钮状态
function updateButtons() {
    elements.convertBtn.disabled = !state.isValidJson || state.converting;
    elements.convertBtn.textContent = state.converting ? '转换中...' : '转换';

    if (state.converting) {
        elements.convertBtn.classList.add('loading');
    } else {
        elements.convertBtn.classList.remove('loading');
    }

    elements.copyBtn.disabled = !state.outputCode;
    elements.downloadBtn.disabled = !state.outputCode;
}

// 从输入元素更新状态
function updateStateFromInputs() {
    state.language = elements.language.value;
    state.structName = elements.structName.value;
    state.packageName = elements.packageName.value;
    state.addComments = elements.addComments.checked;
    state.usePointer = elements.usePointer.checked;
    state.timeAsString = elements.timeAsString.checked;
}

// 更新标签部分可见性
function updateTagSection() {
    elements.tagSection.style.display = state.language === 'golang' ? 'block' : 'none';
}

// 更新包名输入框可见性
function updatePackageGroup() {
    elements.packageGroup.style.display = state.language === 'golang' ? 'block' : 'none';
}

// 切换标签
function toggleTag(tagValue) {
    const index = state.tags.indexOf(tagValue);
    if (index === -1) {
        state.tags.push(tagValue);
    } else {
        state.tags.splice(index, 1);
    }
    renderTags();
}

// 加载示例
function loadExample(example) {
    state.jsonInput = example.json;
    elements.jsonInput.value = example.json;
    validateJson();
    updateUI();
}

// 转换JSON
async function convertJson() {
    if (!state.isValidJson || state.converting) {
        return;
    }

    state.converting = true;
    updateUI();

    try {
        const response = await fetch('/api/v1/convert', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                json: state.jsonInput,
                language: state.language,
                tags: state.tags,
                struct_name: state.structName,
                package_name: state.packageName,
                make_public: false,
                add_comments: state.addComments,
                use_pointer: state.usePointer,
                time_as_string: state.timeAsString
            })
        });

        const data = await response.json();

        if (data.success) {
            state.outputCode = data.code;
            elements.outputCode.textContent = data.code;

            // 语法高亮
            applySyntaxHighlighting();
        } else {
            state.outputCode = `错误: ${data.error}`;
            elements.outputCode.textContent = `错误: ${data.error}`;
        }
    } catch (error) {
        state.outputCode = `请求失败: ${error.message}`;
        elements.outputCode.textContent = `请求失败: ${error.message}`;
    } finally {
        state.converting = false;
        updateUI();
    }
}

// 应用语法高亮
function applySyntaxHighlighting() {
    const codeElement = elements.outputCode;
    const code = codeElement.textContent;

    // 简单的语法高亮（实际项目中可以使用highlight.js）
    if (state.language === 'golang') {
        codeElement.innerHTML = highlightGoCode(code);
    } else if (state.language === 'java') {
        codeElement.innerHTML = highlightJavaCode(code);
    } else {
        codeElement.innerHTML = escapeHtml(code);
    }
}

// 简单的HTML转义
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

// Go代码高亮（简单版）
function highlightGoCode(code) {
    const keywords = ['package', 'import', 'type', 'struct', 'interface', 'func',
                     'var', 'const', 'map', 'chan', 'go', 'select', 'defer'];

    let highlighted = escapeHtml(code);

    keywords.forEach(keyword => {
        const regex = new RegExp(`\\b${keyword}\\b`, 'g');
        highlighted = highlighted.replace(regex, `<span class="keyword">${keyword}</span>`);
    });

    // 高亮注释
    highlighted = highlighted.replace(/\/\/.*$/gm, '<span class="comment">$&</span>');

    // 高亮字符串
    highlighted = highlighted.replace(/"[^"]*"/g, '<span class="string">$&</span>');

    // 高亮数字
    highlighted = highlighted.replace(/\b\d+\b/g, '<span class="number">$&</span>');

    return highlighted;
}

// Java代码高亮（简单版）
function highlightJavaCode(code) {
    const keywords = ['import', 'package', 'public', 'private', 'protected',
                     'class', 'interface', 'void', 'int', 'long', 'float',
                     'double', 'boolean', 'char', 'byte', 'short', 'new'];

    let highlighted = escapeHtml(code);

    keywords.forEach(keyword => {
        const regex = new RegExp(`\\b${keyword}\\b`, 'g');
        highlighted = highlighted.replace(regex, `<span class="keyword">${keyword}</span>`);
    });

    // 高亮注释
    highlighted = highlighted.replace(/\/\/.*$/gm, '<span class="comment">$&</span>');
    highlighted = highlighted.replace(/\/\*[\s\S]*?\*\//g, '<span class="comment">$&</span>');

    // 高亮字符串
    highlighted = highlighted.replace(/"[^"]*"/g, '<span class="string">$&</span>');

    // 高亮注解
    highlighted = highlighted.replace(/@\w+/g, '<span class="annotation">$&</span>');

    return highlighted;
}

// 复制到剪贴板
async function copyToClipboard() {
    if (!state.outputCode) {
        return;
    }

    try {
        await navigator.clipboard.writeText(state.outputCode);
        alert('代码已复制到剪贴板！');
    } catch (error) {
        console.error('复制失败:', error);
        alert('复制失败，请手动复制');
    }
}

// 下载代码
function downloadCode() {
    if (!state.outputCode) {
        return;
    }

    const extension = getFileExtension();
    const filename = `${state.structName}.${extension}`;
    const blob = new Blob([state.outputCode], { type: 'text/plain; charset=utf-8' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = filename;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
}

// 获取文件扩展名
function getFileExtension() {
    switch (state.language) {
        case 'golang': return 'go';
        case 'java': return 'java';
        case 'typescript': return 'ts';
        case 'python': return 'py';
        default: return 'txt';
    }
}

// 清空输出
function clearOutput() {
    state.outputCode = '';
    elements.outputCode.textContent = '';
    updateUI();
}

// 页面加载完成时初始化应用
document.addEventListener('DOMContentLoaded', initApp);

// 导出全局函数（用于示例点击）
window.loadExample = loadExample;