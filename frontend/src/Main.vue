<template >
    <a-row class="no-scrollbar">
        <a-col :span="4">
            <a-menu class="scroll-container-1" v-model:selectedKeys="store.selectedKeys" :open-keys="store.openKeys" 
                @openChange="store.onOpenChange" mode="inline">
                <a-menu-item key="index">
                    <template #icon>
                        <home-outlined />
                    </template>
                    {{ menuRecord['index'] }}
                </a-menu-item>
                <a-sub-menu key="page_edit">
                    <template #title>编辑功能</template>
                    <template #icon>
                        <form-outlined />
                    </template>
                    <a-menu-item key="merge">
                        <template #icon>
                            <merge-cells-outlined />
                        </template>
                        {{ menuRecord['merge'] }}
                    </a-menu-item>
                    <a-menu-item key="split">
                        <template #icon>
                            <split-cells-outlined />
                        </template>
                        {{ menuRecord['split'] }}
                    </a-menu-item>
                    <a-menu-item key="insert">
                        <template #icon>
                            <login-outlined />
                        </template>
                        {{ menuRecord['insert'] }}
                    </a-menu-item>
                    <a-menu-item key="rotate">
                        <template #icon>
                            <rotate-right-outlined />
                        </template>
                        {{ menuRecord['rotate'] }}
                    </a-menu-item>
                    <a-menu-item key="delete">
                        <template #icon>
                            <delete-outlined />
                        </template>
                        {{ menuRecord['delete'] }}
                    </a-menu-item>
                    <a-menu-item key="reorder">
                        <template #icon>
                            <ordered-list-outlined />
                        </template>
                        {{ menuRecord['reorder'] }}
                    </a-menu-item>
                    <a-menu-item key="extract">
                        <template #icon>
                            <aim-outlined />
                        </template>
                        {{ menuRecord['extract'] }}
                    </a-menu-item>
                    <a-menu-item key="convert">
                        <template #icon>
                            <sync-outlined />
                        </template>
                        {{ menuRecord['convert'] }}
                    </a-menu-item>
                    <a-menu-item key="crop">
                        <template #icon>
                            <scissor-outlined />
                        </template>
                        {{ menuRecord['crop'] }}
                    </a-menu-item>
                    <a-menu-item key="scale">
                        <template #icon>
                            <fullscreen-outlined />
                        </template>
                        {{ menuRecord['scale'] }}
                    </a-menu-item>
                    <a-menu-item key="cut">
                        <template #icon>
                            <borderless-table-outlined />
                        </template>
                        {{ menuRecord['cut'] }}
                    </a-menu-item>
                    <a-menu-item key="header">
                        <template #icon>
                            <credit-card-outlined />
                        </template>
                        {{ menuRecord['header'] }}
                    </a-menu-item>
                    <a-menu-item key="page_number">
                        <template #icon>
                            <field-binary-outlined />
                        </template>
                        {{ menuRecord['page_number'] }}
                    </a-menu-item>
                </a-sub-menu>
                <a-sub-menu key="protect">
                    <template #title>保护功能</template>
                    <template #icon>
                        <safety-certificate-outlined />
                    </template>
                    <a-menu-item key="watermark">
                        <template #icon>
                            <highlight-outlined />
                        </template>
                        {{ menuRecord['watermark'] }}
                    </a-menu-item>
                    <a-menu-item key="encrypt">
                        <template #icon>
                            <lock-outlined />
                        </template>
                        {{ menuRecord['encrypt'] }}
                    </a-menu-item>
                    <a-menu-item key="sign">
                        <template #icon>
                            <edit-outlined />
                        </template>
                        {{ menuRecord['sign'] }}
                    </a-menu-item>
                </a-sub-menu>
                <a-sub-menu key="other">
                    <template #title>其他功能</template>
                    <template #icon>
                        <appstore-outlined />
                    </template>
                    <!-- <a-menu-item key="meta">
                        <template #icon>
                            <info-circle-outlined />
                        </template>
                        {{ menuRecord['meta'] }}
                    </a-menu-item> -->
                    <a-menu-item key="bookmark">
                        <template #icon>
                            <book-outlined />
                        </template>
                        {{ menuRecord['bookmark'] }}
                    </a-menu-item>
                    <a-menu-item key="compress">
                        <template #icon>
                            <file-zip-outlined />
                        </template>
                        {{ menuRecord['compress'] }}
                    </a-menu-item>
                    <a-menu-item key="ocr">
                        <template #icon>
                            <eye-outlined />
                        </template>
                        {{ menuRecord['ocr'] }}
                    </a-menu-item>
                    <a-menu-item key="background">
                        <template #icon>
                            <bg-colors-outlined />
                        </template>
                        {{ menuRecord['background'] }}
                    </a-menu-item>
                    <a-menu-item key="annot">
                        <template #icon>
                            <message-outlined />
                        </template>
                        {{ menuRecord['annot'] }}
                    </a-menu-item>
                    <!-- <a-menu-item key="crack">
                        <template #icon>
                            <tool-outlined />
                        </template>
                        {{ menuRecord['crack'] }}
                    </a-menu-item> -->
                    <a-menu-item key="dual">
                        <template #icon>
                            <file-search-outlined />
                        </template>
                        {{ menuRecord['dual'] }}
                    </a-menu-item>
                </a-sub-menu>
                <a-menu-item key="settings">
                    <template #icon>
                        <SettingOutlined />
                    </template>
                    {{ menuRecord['settings'] }}
                </a-menu-item>
                <!-- <a-menu-item key="debug">
                    <template #icon>
                        <SettingOutlined />
                    </template>
                    {{ menuRecord['debug'] }}
                </a-menu-item> -->
            </a-menu>
        </a-col>
        <a-col :span="20">
            <div style="margin-left: 1vw">
                <div style="margin-right: 5vw;margin-top: 0.5em; " v-if="store.selectedKeys.at(0) !== 'index'">
                    <a-typography-title>{{ menuRecord[store.selectedKeys.at(0) || "merge"] }}</a-typography-title>
                    <a-typography-paragraph>
                        <blockquote>功能说明：{{ menuDesc[store.selectedKeys.at(0) || "merge"] }}</blockquote>
                    </a-typography-paragraph>
                </div>
                <div style="margin-right: 5vw;margin-top: 0.5em;" v-else>
                    <a-typography-title>功能列表</a-typography-title>
                </div>
                <div class="scroll-container-2">
                    <Index v-if="store.selectedKeys.at(0) === 'index'" />
                    <MergeForm v-if="store.selectedKeys.at(0) === 'merge'" />
                    <SplitForm v-if="store.selectedKeys.at(0) === 'split'" />
                    <InsertForm v-if="store.selectedKeys.at(0) === 'insert'" />
                    <RotateForm v-if="store.selectedKeys.at(0) === 'rotate'" />
                    <DeleteForm v-if="store.selectedKeys.at(0) === 'delete'" />
                    <ReorderForm v-if="store.selectedKeys.at(0) === 'reorder'" />
                    <BookmarkForm v-if="store.selectedKeys.at(0) === 'bookmark'" />
                    <ScaleForm v-if="store.selectedKeys.at(0) === 'scale'" />
                    <WatermarkForm v-if="store.selectedKeys.at(0) === 'watermark'" />
                    <CropForm v-if="store.selectedKeys.at(0) === 'crop'" />
                    <CutForm v-if="store.selectedKeys.at(0) === 'cut'" />
                    <ExtractForm v-if="store.selectedKeys.at(0) === 'extract'" />
                    <CompressForm v-if="store.selectedKeys.at(0) === 'compress'" />
                    <ConvertForm v-if="store.selectedKeys.at(0) === 'convert'" />
                    <EncryptForm v-if="store.selectedKeys.at(0) === 'encrypt'" />
                    <OcrForm v-if="store.selectedKeys.at(0) === 'ocr'" />
                    <PreferencesForm v-if="store.selectedKeys.at(0) === 'settings'" />
                    <HeaderAndFooterForm v-if="store.selectedKeys.at(0) === 'header'" />
                    <PageNumberForm v-if="store.selectedKeys.at(0) === 'page_number'" />
                    <BackgroundForm v-if="store.selectedKeys.at(0) === 'background'" />
                    <MetaForm v-if="store.selectedKeys.at(0) === 'meta'" />
                    <DualLayerForm v-if="store.selectedKeys.at(0) === 'dual'" />
                    <PasswordCrackForm v-if="store.selectedKeys.at(0) === 'crack'" />
                    <SignForm v-if="store.selectedKeys.at(0) === 'sign'" />
                    <AnnotForm v-if="store.selectedKeys.at(0) === 'annot'" />
                    <Debug v-if="store.selectedKeys.at(0) === 'debug'" />
                </div>
            </div>
        </a-col>
    </a-row>
</template>
<script lang="ts">
import { defineComponent, reactive, watch, ref } from 'vue';
import {
    MinusCircleOutlined,
    PlusOutlined,
    AppstoreOutlined,
    SettingOutlined,
    BarsOutlined,
    BorderlessTableOutlined,
    ScissorOutlined,
    CopyOutlined,
    BlockOutlined,
    FileZipOutlined,
    RotateRightOutlined,
    LockOutlined,
    OrderedListOutlined,
    HighlightOutlined,
    AimOutlined,
    ExportOutlined,
    UploadOutlined,
    FullscreenOutlined,
    CloseOutlined,
    FontColorsOutlined,
    FontSizeOutlined,
    EyeOutlined,
    SplitCellsOutlined,
    LoginOutlined,
    FormOutlined,
    DeleteOutlined,
    MergeCellsOutlined,
    InfoCircleOutlined,
    BookOutlined,
    CreditCardOutlined,
    BgColorsOutlined,
    SafetyCertificateOutlined,
    TabletOutlined,
    FieldBinaryOutlined,
    SyncOutlined,
    FileSearchOutlined,
    ToolOutlined,
    HomeOutlined,
    EditOutlined,
    StarOutlined,
    MessageOutlined,
    createFromIconfontCN,
} from '@ant-design/icons-vue';

import { menuDesc, menuRecord } from "./components/data";
import MergeForm from "./components/Forms/MergeForm.vue";
import SplitForm from "./components/Forms/SplitForm.vue";
import DeleteForm from "./components/Forms/DeleteForm.vue";
import ReorderForm from "./components/Forms/ReorderForm.vue";
import InsertForm from "./components/Forms/InsertForm.vue";
import BookmarkForm from "./components/Forms/BookmarkForm.vue";
import ScaleForm from "./components/Forms/ScaleForm.vue";
import WatermarkForm from './components/Forms/WatermarkForm.vue';
import RotateForm from "./components/Forms/RotateForm.vue";
import CropForm from "./components/Forms/CropForm.vue";
import CutForm from "./components/Forms/CutForm.vue";
import ExtractForm from "./components/Forms/ExtractForm.vue";
import CompressForm from "./components/Forms/CompressForm.vue";
import ConvertForm from "./components/Forms/ConvertForm.vue";
import EncryptForm from "./components/Forms/EncryptForm.vue";
import OcrForm from "./components/Forms/OcrForm.vue";
import PreferencesForm from "./components/Forms/PreferencesForm.vue";
import HeaderAndFooterForm from "./components/Forms/HeaderAndFooterForm.vue";
import BackgroundForm from "./components/Forms/BackgroundForm.vue";
import PageNumberForm from "./components/Forms/PageNumberForm.vue";
import MetaForm from "./components/Forms/MetaForm.vue";
import DualLayerForm from "./components/Forms/DualLayerForm.vue";
import PasswordCrackForm from "./components/Forms/PasswordCrackForm.vue";
import Index from "./components/Forms/Index.vue";
import SignForm from "./components/Forms/SignForm.vue";
import AnnotForm from "./components/Forms/AnnotForm.vue";
import Debug from "./components/Forms/Debug.vue";
import { useMenuState } from './store/menu';

const IconFont = createFromIconfontCN({ scriptUrl: '//at.alicdn.com/t/font_8d5l8fzk5b87iudi.js' });

export default defineComponent({
    components: {
        // icon
        MinusCircleOutlined,
        PlusOutlined,
        AppstoreOutlined,
        SettingOutlined,
        BarsOutlined,
        BorderlessTableOutlined,
        ScissorOutlined,
        CopyOutlined,
        BlockOutlined,
        FileZipOutlined,
        RotateRightOutlined,
        LockOutlined,
        OrderedListOutlined,
        HighlightOutlined,
        AimOutlined,
        ExportOutlined,
        UploadOutlined,
        FullscreenOutlined,
        CloseOutlined,
        FontColorsOutlined,
        FontSizeOutlined,
        EyeOutlined,
        SplitCellsOutlined,
        LoginOutlined,
        FormOutlined,
        DeleteOutlined,
        MergeCellsOutlined,
        InfoCircleOutlined,
        BookOutlined,
        CreditCardOutlined,
        BgColorsOutlined,
        SafetyCertificateOutlined,
        TabletOutlined,
        FieldBinaryOutlined,
        SyncOutlined,
        FileSearchOutlined,
        ToolOutlined,
        HomeOutlined,
        EditOutlined,
        StarOutlined,
        MessageOutlined,
        IconFont,
        // form
        // 合并
        MergeForm,
        // 拆分
        SplitForm,
        // 删除
        DeleteForm,
        // 重排
        ReorderForm,
        // 插入/替换
        InsertForm,
        // 书签
        BookmarkForm,
        // 缩放
        ScaleForm,
        // 水印
        WatermarkForm,
        // 旋转
        RotateForm,
        // 裁剪
        CropForm,
        // 分割/组合
        CutForm,
        // 提取
        ExtractForm,
        // 压缩
        CompressForm,
        // 转换
        ConvertForm,
        // 加密
        EncryptForm,
        // OCR
        OcrForm,
        // 首选项
        PreferencesForm,
        // 页眉页脚
        HeaderAndFooterForm,
        // 文档背景
        BackgroundForm,
        // 页码设置
        PageNumberForm,
        // 文档属性
        MetaForm,
        // 双层PDF
        DualLayerForm,
        // 密码破解
        PasswordCrackForm,
        // 首页
        Index,
        // 电子签名
        SignForm,
        // 注释
        AnnotForm,
        // 调试
        Debug,
    },
    setup() {
        const store = useMenuState();
        return {
            menuRecord,
            menuDesc,
            store,
        };
    },
});
</script>

<style>
.scroll-container-1 {
  height: 100vh; /* 设置一个固定高度 */
  overflow-y: auto
}

.scroll-container-2 {
  height: 80vh; /* 设置一个固定高度 */
  overflow-y: auto; /* 添加垂直滚动条*/
}
.no-scrollbar {
  overflow-y: hidden;  /*隐藏竖直滚动条 */
  height: 100vh    /*设置高度为视口高度，确保内容不会超出视口 */
}
</style>