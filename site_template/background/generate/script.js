Array.prototype.contains = function (element) {
    return this.indexOf(element) > -1;
};

const VueClipboard = window['VueClipboard'];
Vue.use(VueClipboard);

// 通过builder内的数据，生成mark
new Vue({
    // root node
    el: "#app",
    // the instance state
    data: {
        data_re: {
            content: {
                value: "",
                show: true,
            },
            link: {
                value: "",
                show: true,
            },
            list: {
                value: "",
                show: true,
            },
        },

        expand: {
            backend_section_label: {
                name: "backend_section_label",
                value: "文章信息",
                content: true,
                link: true,
                show: true,
            },
            backend_env_label: {
                name: "backend_env_label",
                value: "所有标签",
                content: true,
                link: true,
                show: true,
            },
            backend_env: {
                name: "backend_env",
                value: "mingcheng",
                content: true,
                link: true,
                show: true,
            },
            backend_section: {
                name: "backend_section",
                value: "shezhi",
                content: true,
                link: true,
                show: true,
            },
            connection_type: {
                value: "tags",
                name: "标签",
            },
            keys: []
        },

        mark: {
            key: {
                name: "key",
                value: "",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            content_id: {
                name: "content_id",
                value: "123",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            backend_section: {
                name: "backend_section",
                value: "shezhi",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            backend_section_label: {
                name: "backend_section_label",
                value: "公司信息",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            backend_env: {
                name: "backend_env",
                value: "mingcheng",
                content: false,
                check: false,
                link: false,
                show: false,
            },
            backend_env_label: {
                name: "backend_env_label",
                value: "名称",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_page: {
                value: "1",
                name: "mark_page",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_page_size: {
                value: "1",
                name: "mark_page_size",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_content_id: {
                value: "1324",
                name: "mark_content_id",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_created_at: {
                value: "2020-06-23",
                name: "mark_created_at",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_updated_at: {
                value: "2020-06-23",
                name: "mark_updated_at",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_title: {
                value: "新华社消息:商务部",
                name: "mark_title",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_type: {
                value: "blog",
                name: "mark_type",
                check: true,
                content: true,
                link: false,
                show: false,
            },
            multiple: {
                value: "false",
                name: "单标记",
                check: true,
                content: true,
                link: false,
                show: false,
            },
            mark_show_on_dashboard: {
                value: "true",
                name: "后台显示",
                check: true,
                content: true,
                link: false,
                show: false,
            },
            mark_abstract: {
                value: "有人说:灾难",
                name: "mark_abstract",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_body: {
                value: "最新消息:灾难",
                name: "mark_body",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_content_url: {
                value: "2020-06-23",
                name: "mark_content_url",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_source_url: {
                value: "/detail.html",
                name: "mark_source_url",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_list_url: {
                value: "/list.html",
                name: "mark_list_url",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_cover: {
                value: "2020-06-23",
                name: "mark_cover",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_size: {
                value: "13MB",
                name: "mark_size",
                check: false,
                content: false,
                link: false,
                show: false,
            },
            mark_copyright_stat: {
                value: "2020-06-23",
                name: "mark_copyright_stat",
                check: false,
                content: false,
                link: false,
                show: false,
            },
        },

        filters: [],
        history: [],
        current: "blog",
    },
    methods: {
        addNewExpand: function () {
            this.expand.backend_section.value = pinyinUtil.getPinyin(this.expand.backend_section_label.value).replace(/\s*/g, "");
            this.expand.backend_env.value = pinyinUtil.getPinyin(this.expand.backend_env_label.value).replace(/\s*/g, "");

            let data_re = 'data-re="backend_section=' + this.expand.backend_section.value + '&backend_env=' + this.expand.backend_env.value + '&mark_type=' + this.expand.connection_type.value + '"'

            this.expand.keys.push({
                backend_section: {
                    name: "backend_section",
                    value: this.expand.backend_section.value,
                    content: true,
                    link: true,
                    show: true,
                },
                backend_section_label: {
                    name: "backend_section_label",
                    value: this.expand.backend_section_label.value,
                    content: true,
                    link: true,
                    show: true,
                },
                backend_env: {
                    name: "backend_env",
                    value: this.expand.backend_env.value,
                    content: true,
                    link: true,
                    show: true,
                },
                backend_env_label: {
                    name: "backend_env_label",
                    value: this.expand.backend_env_label.value,
                    content: true,
                    link: true,
                    show: true,
                },
                connection_type: this.expand.connection_type,
                data_re: data_re,
            })

            this.JoinToCode()
        },
        onCopy: function (e) {
            this.$snotify.success('复制成功');
        },
        onError: function (e) {
        },
        JoinToCode: function () {
            for (let key of Object.keys(this.expand.keys)) {
                this.expand.keys[key].backend_section.value = pinyinUtil.getPinyin(this.expand.keys[key].backend_section_label.value).replace(/\s*/g, "");
                this.expand.keys[key].backend_env.value = pinyinUtil.getPinyin(this.expand.keys[key].backend_env_label.value).replace(/\s*/g, "");
                this.expand.keys[key].data_re = 'data-re="backend_section=' + this.expand.keys[key].backend_section.value + '&backend_env=' + this.expand.keys[key].backend_env.value + '&mark_type=' + this.expand.keys[key].connection_type.value + '"';
            }

            this.mark.content_id.value = this.mark.mark_content_id.value

            if (this.mark.mark_type.value !== this.current) {
                this.InitEnv()
                this.current = this.mark.mark_type.value
            }

            this.mark.backend_section.value = pinyinUtil.getPinyin(this.mark.backend_section_label.value).replace(/\s*/g, "");
            this.mark.backend_env.value = pinyinUtil.getPinyin(this.mark.backend_env_label.value).replace(/\s*/g, "");
            this.mark.key.value = this.mark.backend_section.value + "#" + this.mark.backend_env.value

            for (let key of Object.keys(this.mark)) {
                this.mark[key].show = !!this.filters[this.mark.mark_type.value].show.contains(key);
            }

            let mark = this.mark
            let keys = []
            this.filters[this.mark.mark_type.value].content.forEach(function (key) {
                if (mark[key].check === true) {
                    keys.push(key + "=" + mark[key].value)
                }
            });

            let many2many = []
            this.expand.keys.forEach(function (key) {
                many2many.push("connections=" + key.connection_type.value + "#" + key.backend_section.value + "#" + key.backend_env.value)
            });

            if (many2many.length !== 0) {
                keys.push(many2many.join("&"))
            }

            this.data_re.content.value = [`data-re="`, keys.join("&"), `"`].join("")

            keys = []
            this.filters[this.mark.mark_type.value].href.forEach(function (key) {
                if (mark[key].check === true) {
                    keys.push(key + "=" + mark[key].value)
                }
            });
            this.data_re.link.value = [`href="`, keys.join("&"), `"`].join("").replace("mark_source_url=", "")
            this.data_re.link.value = this.data_re.link.value.replace("&", "?")

            keys = []
            this.filters[this.mark.mark_type.value].list.forEach(function (key) {
                console.log(key)
                if (mark[key].check === true) {
                    keys.push(key + "=" + mark[key].value)
                }
            });

            this.data_re.list.value = [`href="`, keys.join("&"), `"`].join("").replace("mark_source_url=", "")
            this.data_re.list.value = this.data_re.list.value.replace("mark_list_url=", "")
            this.data_re.list.value = this.data_re.list.value.replace("&", "?")
        },
        InitEnv() {
            this.filters['blog'] = {
                check: ["key", "mark_show_on_dashboard", "multiple", "mark_type", "mark_title", "mark_list_url", "mark_content_id", "mark_source_url", "backend_section_label", "backend_env_label", "backend_section", "content_id", "backend_env", "mark_page", "mark_page_size"],

                show: [
                    "mark_type",
                    "multiple",
                    "mark_show_on_dashboard",
                    "mark_content_id", "mark_created_at",
                    "mark_title", "mark_abstract", "mark_source_url",
                    "mark_list_url",
                    "mark_source_url", "mark_body", "backend_section_label", "backend_env_label", "mark_page", "mark_page_size"],

                content: ["key", "multiple", "mark_show_on_dashboard", "mark_content_id", "mark_created_at",
                    "mark_title", "mark_abstract", "mark_body", "mark_type", "backend_section_label", "backend_env_label", "backend_section", "backend_env", "mark_page",
                    "mark_page_size"],

                href: ["mark_source_url", "content_id"],
                list: ["mark_list_url", "backend_env", "backend_section", "key"],

                listShow: true,
                linkShow: true,

                contentShow: true,
                thumbnails: true,
            }

            for (let key of Object.keys(this.mark)) {
                this.mark[key].show = !!this.filters[this.mark.mark_type.value].show.contains(key);
                this.mark[key].content = !!this.filters[this.mark.mark_type.value].content.contains(key);
                this.mark[key].check = !!this.filters[this.mark.mark_type.value].check.contains(key);
                this.mark[key].href = !!this.filters[this.mark.mark_type.value].href.contains(key);
                this.mark[key].list = !!this.filters[this.mark.mark_type.value].list.contains(key);
            }

            this.data_re.list.show = this.filters[this.mark.mark_type.value].listShow
            this.data_re.link.show = this.filters[this.mark.mark_type.value].linkShow
            this.data_re.content.show = this.filters[this.mark.mark_type.value].contentShow
        },
    },
    created() {
        this.InitEnv()
        this.JoinToCode()
    },
    computed: {
        listenChange() {
            const {
                mark, expand
            } = this
            return {
                mark, expand
            }
        }
    },
    watch: {
        listenChange: {
            handler: function (n, o) {
                this.JoinToCode()
            },
            deep: true
        }
    }
});
