<template>
    <div class="topo-attribute-wrapper" v-show="isShow">
        <div class="loading" v-bkloading="{isLoading: attr.isLoading}">
            <i class="bk-icon icon-close" @click="closePop"></i>
            <div class="attr-title">{{objName}} {{instName}}</div>
            <div class="attribute-box" id="box">
                <v-attribute 
                    id="attribute"
                    ref="attribute"
                    :formFields="attr.formFields"
                    :formValues="attr.formValues"
                    :type="attr.type"
                    :showBtnGroup="false"
                    :active="isShow"
                    :objId="objId">
                </v-attribute>
            </div>
        </div>
    </div>
</template>

<script>
    import vAttribute from '@/components/object/attribute'
    import { mapGetters } from 'vuex'
    export default {
        props: {
            isShow: {
                type: Boolean,
                default: false
            },
            objId: {
                type: String,
                default: ''
            },
            instId: {
                default: ''
            },
            objName: {
                type: String,
                default: ''
            },
            instName: {
                type: String,
                default: ''
            }
        },
        data () {
            return {
                attr: {
                    formFields: [],
                    formValues: {},
                    type: 'update',
                    isLoading: false
                }
            }
        },
        computed: {
            ...mapGetters([
                'bkSupplierAccount'
            ]),
            ...mapGetters('object', [
                'attribute'
            ]),
            formValuesConfig () {
                let config = {
                    url: '',
                    params: {
                        page: {},
                        fields: {},
                        condition: {}
                    }
                }
                if (this.objId === 'biz') {
                    config.url = `biz/search/${this.bkSupplierAccount}`
                    config.params.condition[this.objId] = [{
                        field: 'bk_biz_id',
                        operator: '$eq',
                        value: this.instId
                    }]
                } else if (this.objId === 'host') {
                    config.url = 'hosts/search'
                    config.params = {
                        page: {
                            start: 0,
                            limit: 10,
                            sort: 'bk_host_id'
                        },
                        pattern: '',
                        bk_biz_id: -1,
                        ip: {
                            flag: 'bk_host_innerip|bk_host_outerip',
                            exact: 1,
                            data: []
                        },
                        condition: [{
                            bk_obj_id: 'host',
                            fields: [],
                            condition: [{
                                field: 'bk_host_id',
                                operator: '$eq',
                                value: this.instId
                            }]
                        }, {
                            bk_obj_id: 'module',
                            fields: [],
                            condition: []
                        }, {
                            bk_obj_id: 'set',
                            fields: [],
                            condition: []
                        }, {
                            bk_obj_id: 'biz',
                            fields: [],
                            condition: [{
                                field: 'default',
                                operator: '$ne',
                                value: 1
                            }]
                        }]
                    }
                } else {
                    config.url = `inst/association/search/owner/${this.bkSupplierAccount}/object/${this.objId}`
                    config.params.condition[this.objId] = [{
                        field: 'bk_inst_id',
                        operator: '$eq',
                        value: this.instId
                    }]
                }
                return config
            }
        },
        watch: {
            isShow (isShow) {
                if (isShow) {
                    this.initData()
                }
            }
        },
        methods: {
            closePop () {
                this.$emit('update:isShow', false)
            },
            resetAttributeBox () {
                let box = document.getElementById('box')
                let attribute = document.getElementById('attribute')
                let topo = document.getElementById('topo')
                if (topo.offsetHeight < box.offsetHeight * 0.8) {
                    box.style.height = `${topo.offsetHeight * 0.8 - 102}px`
                } else if (box.offsetHeight > attribute.offsetHeight) {
                    box.style.height = `${attribute.offsetHeight + 40}px`
                } else {
                    box.style.height = `${topo.offsetHeight * 0.8 - 102}px`
                }
            },
            async initData () {
                this.attr.isLoading = true
                await Promise.all([
                    this.$store.dispatch('object/getAttribute', this.objId),
                    this.getFormValues()
                ])
                this.attr.formFields = this.attribute[this.objId]
                this.attr.isLoading = false
                this.$nextTick(() => {
                    this.resetAttributeBox()
                })
            },
            async getFormValues () {
                try {
                    let res = await this.$axios.post(this.formValuesConfig.url, this.formValuesConfig.params)
                    this.attr.formValues = this.objId === 'host' ? res.data.info[0].host : res.data.info[0]
                } catch (e) {
                    this.$alertMsg(e.message || e.data['bk_error_msg'] || e.statusText)
                }
            }
        },
        created () {
            window.onresize = () => {
                this.resetAttributeBox()
            }
        },
        components: {
            vAttribute
        }
    }
</script>

<style lang="scss" scoped>
    .topo-attribute-wrapper {
        position: absolute;
        background: #fff;
        width: 710px;
        max-height: 80%;
        min-height: 200px;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        box-shadow: 0px 2px 9.6px 0.4px rgba(0, 0, 0, 0.4);
        .loading {
            min-height: 200px;
            height: 100%;
        }
        .attr-title {
            font-size: 16px;
            padding: 40px 32px 0;
            color: #737987;
        }
        .icon-close {
            padding: 2px;
            font-size: 16px;
            position: absolute;
            right: 10px;
            top: 10px;
            cursor: pointer;
        }
        .attribute-box {
            height: calc(100% - 40px);
            overflow: auto;
            @include scrollbar;
            padding-bottom: 40px;
            margin: 20px 0;
        }
    }
</style>

