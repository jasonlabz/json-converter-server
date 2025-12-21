// Package static
//
//   _ __ ___   __ _ _ __  _   _| |_
//  | '_ ` _ \ / _` | '_ \| | | | __|
//  | | | | | | (_| | | | | |_| | |_
//  |_| |_| |_|\__,_|_| |_|\__,_|\__|
//
//  Buddha bless, no bugs forever!
//
//  Author:    lucas
//  Email:     1783022886@qq.com
//  Created:   2025/12/21 22:20
//  Version:   v1.0.0

package static

const INDEX_HTML = `<!-- frontend/index.html -->
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>JSON转结构体工具</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css" rel="stylesheet">
    <style>
        .editor-container {
            height: 400px;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        .code-container {
            background-color: #f5f5f5;
            border: 1px solid #ddd;
            border-radius: 4px;
            padding: 15px;
            max-height: 500px;
            overflow-y: auto;
        }
        .tag-badge {
            margin-right: 5px;
            margin-bottom: 5px;
            cursor: pointer;
        }
        .tag-badge.active {
            background-color: #0d6efd;
        }
    </style>
</head>
<body>
    <div id="app" class="container-fluid">
        <div class="row mt-4">
            <div class="col-12">
                <h1 class="text-center">
                    <i class="fas fa-code me-2"></i>JSON转结构体工具
                </h1>
                <p class="text-center text-muted">
                    输入JSON数据，自动生成多种语言的结构体定义
                </p>
            </div>
        </div>

        <div class="row mt-4">
            <!-- 左侧：JSON输入 -->
            <div class="col-md-6">
                <div class="card">
                    <div class="card-header bg-primary text-white">
                        <i class="fas fa-keyboard me-2"></i>JSON输入
                    </div>
                    <div class="card-body">
                        <div class="mb-3">
                            <label class="form-label">JSON数据</label>
                            <textarea 
                                class="form-control" 
                                rows="15"
                                v-model="inputJson"
                                placeholder='{"name": "John", "age": 30, "email": "john@example.com"}'
                                @input="validateJson"
                            ></textarea>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">结构体名称</label>
                            <input type="text" class="form-control" v-model="structName">
                        </div>
                    </div>
                </div>
            </div>

            <!-- 右侧：配置选项 -->
            <div class="col-md-6">
                <div class="card">
                    <div class="card-header bg-success text-white">
                        <i class="fas fa-cog me-2"></i>配置选项
                    </div>
                    <div class="card-body">
                        <!-- 语言选择 -->
                        <div class="mb-3">
                            <label class="form-label">目标语言</label>
                            <select class="form-select" v-model="language">
                                <option v-for="lang in languages" :value="lang.value">
                                    {{ lang.label }}
                                </option>
                            </select>
                        </div>

                        <!-- 标签选择 -->
                        <div class="mb-3" v-if="language === 'golang'">
                            <label class="form-label">标签类型</label>
                            <div>
                                <span 
                                    v-for="tag in availableTags" 
                                    :key="tag.value"
                                    class="badge tag-badge"
                                    :class="{'bg-primary': selectedTags.includes(tag.value), 'bg-secondary': !selectedTags.includes(tag.value)}"
                                    @click="toggleTag(tag.value)"
                                >
                                    {{ tag.label }}
                                </span>
                            </div>
                        </div>

                        <!-- Go特有选项 -->
                        <div class="row" v-if="language === 'golang'">
                            <div class="col-md-6">
                                <div class="mb-3">
                                    <label class="form-label">包名</label>
                                    <input type="text" class="form-control" v-model="packageName">
                                </div>
                            </div>
                            <div class="col-md-6">
                                <div class="mb-3">
                                    <label class="form-label">
                                        <input type="checkbox" class="form-check-input" v-model="usePointer">
                                        使用指针类型
                                    </label>
                                </div>
                            </div>
                        </div>

                        <!-- Java特有选项 -->
                        <div v-if="language === 'java'">
                            <div class="mb-3">
                                <label class="form-label">
                                    <input type="checkbox" class="form-check-input" v-model="makePublic">
                                    生成public字段
                                </label>
                            </div>
                        </div>

                        <!-- 通用选项 -->
                        <div class="row">
                            <div class="col-md-6">
                                <div class="mb-3">
                                    <label class="form-label">
                                        <input type="checkbox" class="form-check-input" v-model="addComments">
                                        添加注释
                                    </label>
                                </div>
                            </div>
                            <div class="col-md-6">
                                <div class="mb-3">
                                    <label class="form-label">
                                        <input type="checkbox" class="form-check-input" v-model="timeAsString">
                                        时间作为字符串
                                    </label>
                                </div>
                            </div>
                        </div>

                        <!-- 转换按钮 -->
                        <div class="d-grid">
                            <button 
                                class="btn btn-primary btn-lg"
                                @click="convertJson"
                                :disabled="!isValidJson || converting"
                            >
                                <i class="fas fa-sync-alt me-2" :class="{'fa-spin': converting}"></i>
                                {{ converting ? '转换中...' : '转换' }}
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- 输出结果 -->
        <div class="row mt-4" v-if="outputCode">
            <div class="col-12">
                <div class="card">
                    <div class="card-header bg-info text-white d-flex justify-content-between align-items-center">
                        <span>
                            <i class="fas fa-code me-2"></i>生成结果
                        </span>
                        <button class="btn btn-sm btn-light" @click="copyToClipboard">
                            <i class="fas fa-copy me-1"></i>复制代码
                        </button>
                    </div>
                    <div class="card-body">
                        <pre class="code-container"><code>{{ outputCode }}</code></pre>
                    </div>
                </div>
            </div>
        </div>

        <!-- 示例JSON -->
        <div class="row mt-4">
            <div class="col-12">
                <div class="card">
                    <div class="card-header">
                        <i class="fas fa-lightbulb me-2"></i>示例JSON
                    </div>
                    <div class="card-body">
                        <div class="row">
                            <div class="col-md-4" v-for="(example, index) in examples" :key="index">
                                <div class="card example-card" @click="loadExample(example.json)">
                                    <div class="card-body">
                                        <h6 class="card-title">{{ example.name }}</h6>
                                        <p class="card-text small text-muted">{{ example.description }}</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script src="https://unpkg.com/vue@3/dist/vue.global.js"></script>
    <script>
        const { createApp, ref, computed, onMounted } = Vue;

        createApp({
            setup() {
                const inputJson = ref('');
                const outputCode = ref('');
                const language = ref('golang');
                const selectedTags = ref(['json', 'mapstructure']);
                const structName = ref('User');
                const packageName = ref('models');
                const makePublic = ref(false);
                const addComments = ref(true);
                const usePointer = ref(false);
                const timeAsString = ref(false);
                const converting = ref(false);
                const isValidJson = ref(true);

                const languages = ref([]);
                const availableTags = ref([]);

                const examples = ref([
                    {
                        name: '用户信息',
                        description: '简单的用户对象',
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
                        description: '包含嵌套对象的产品',
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
                        description: '包含数组的订单信息',
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
                ]);

                // 初始化数据
                const initData = async () => {
                    try {
                        const [langRes, tagRes] = await Promise.all([
                            fetch('/api/v1/languages'),
                            fetch('/api/v1/tags')
                        ]);
                        
                        languages.value = await langRes.json();
                        availableTags.value = await tagRes.json();
                    } catch (error) {
                        console.error('Failed to load data:', error);
                    }
                };

                // JSON验证
                const validateJson = () => {
                    if (!inputJson.value.trim()) {
                        isValidJson.value = true;
                        return;
                    }
                    
                    try {
                        JSON.parse(inputJson.value);
                        isValidJson.value = true;
                    } catch (e) {
                        isValidJson.value = false;
                    }
                };

                // 切换标签
                const toggleTag = (tag) => {
                    const index = selectedTags.value.indexOf(tag);
                    if (index === -1) {
                        selectedTags.value.push(tag);
                    } else {
                        selectedTags.value.splice(index, 1);
                    }
                };

                // 转换JSON
                const convertJson = async () => {
                    if (!inputJson.value.trim()) {
                        alert('请输入JSON数据');
                        return;
                    }

                    converting.value = true;
                    
                    try {
                        const response = await fetch('/api/v1/convert', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify({
                                json: inputJson.value,
                                language: language.value,
                                tags: selectedTags.value,
                                struct_name: structName.value,
                                package_name: packageName.value,
                                make_public: makePublic.value,
                                add_comments: addComments.value,
                                use_pointer: usePointer.value,
                                time_as_string: timeAsString.value
                            })
                        });

                        const data = await response.json();
                        
                        if (data.success) {
                            outputCode.value = data.code;
                        } else {
                            alert('转换失败: ' + data.error);
                        }
                    } catch (error) {
                        alert('请求失败: ' + error.message);
                    } finally {
                        converting.value = false;
                    }
                };

                // 加载示例
                const loadExample = (json) => {
                    inputJson.value = json;
                    validateJson();
                };

                // 复制到剪贴板
                const copyToClipboard = async () => {
                    try {
                        await navigator.clipboard.writeText(outputCode.value);
                        alert('代码已复制到剪贴板！');
                    } catch (err) {
                        console.error('复制失败:', err);
                    }
                };

                onMounted(() => {
                    initData();
                    // 加载第一个示例
                    if (examples.value.length > 0) {
                        loadExample(examples.value[0].json);
                    }
                });

                return {
                    inputJson,
                    outputCode,
                    language,
                    selectedTags,
                    structName,
                    packageName,
                    makePublic,
                    addComments,
                    usePointer,
                    timeAsString,
                    converting,
                    isValidJson,
                    languages,
                    availableTags,
                    examples,
                    validateJson,
                    toggleTag,
                    convertJson,
                    loadExample,
                    copyToClipboard
                };
            }
        }).mount('#app');
    </script>
</body>
</html>`
