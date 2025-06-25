<template>
  <div class="container mt-4">
    <h2>NFT 元数据管理</h2>

    <div class="row">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            创建新元数据
          </div>
          <div class="card-body">
            <form @submit.prevent="createMetadata">
              <div class="mb-3">
                <label for="name" class="form-label">名称</label>
                <input type="text" class="form-control" id="name" v-model="newMetadata.name" required>
              </div>

              <div class="mb-3">
                <label for="description" class="form-label">描述</label>
                <textarea class="form-control" id="description" v-model="newMetadata.description" rows="3"
                  required></textarea>
              </div>

              <div class="mb-3">
                <label for="image" class="form-label">NFT图片</label>
                
                <!-- 图片上传区域 -->
                <div class="image-upload-section">
                  <!-- 上传按钮 -->
                  <div class="mb-2">
                    <input type="file" class="d-none" ref="imageInput" @change="handleImageSelect" 
                      accept="image/jpeg,image/jpg,image/png,image/gif,image/webp,image/svg+xml,image/bmp">
                    <button type="button" class="btn btn-outline-primary btn-sm" @click="$refs.imageInput.click()"
                      :disabled="isUploadingImage">
                      <span v-if="isUploadingImage" class="spinner-border spinner-border-sm me-1"></span>
                      <i v-else class="bi bi-cloud-upload me-1"></i>
                      {{ isUploadingImage ? '上传中...' : '选择图片文件' }}
                    </button>
                    <small class="text-muted ms-2">支持: JPG, PNG, GIF, WebP, SVG, BMP (最大10MB)</small>
                  </div>

                  <!-- 图片预览 -->
                  <div v-if="newMetadata.image" class="mb-2">
                    <div class="image-preview-container">
                      <img :src="getImageDisplayUrl(newMetadata.image)" 
                        alt="NFT图片预览" 
                        class="image-preview"
                        @error="handleImageError">
                      <button type="button" class="btn btn-danger btn-sm position-absolute top-0 end-0 m-1"
                        @click="removeImage" title="移除图片">
                        <i class="bi bi-x"></i>
                      </button>
                    </div>
                  </div>

                  <!-- 手动输入URL -->
                  <div class="mb-2">
                    <input type="url" class="form-control form-control-sm" 
                      v-model="newMetadata.image"
                      placeholder="或手动输入图片URL: https://example.com/image.jpg 或 ipfs://...">
                  </div>

                  <div class="form-text">
                    <i class="bi bi-info-circle me-1"></i>
                    可以上传图片文件或手动输入图片URL。上传的图片将自动存储到IPFS网络。
                  </div>
                </div>
              </div>

              <div class="mb-3">
                <label for="policy" class="form-label">访问策略</label>
                
                <!-- 策略编辑器 -->
                <div class="policy-editor border rounded p-3 bg-light">
                  <div class="d-flex justify-content-between align-items-center mb-3">
                    <h6 class="mb-0">
                      <i class="bi bi-shield-lock me-2"></i>访问控制策略构建器
                    </h6>
                    <button type="button" class="btn btn-sm btn-outline-info" @click="addPolicyGroup">
                      <i class="bi bi-plus-square me-1"></i>新建条件组
                    </button>
                  </div>

                  <!-- 策略组列表 -->
                  <div v-if="policyGroups.length > 0" class="mb-3">
                    <div v-for="(group, groupIndex) in policyGroups" :key="group.id" 
                         class="policy-group mb-3 p-3 border rounded bg-white shadow-sm">
                      
                      <!-- 组标题和逻辑选择 -->
                      <div class="d-flex justify-content-between align-items-center mb-3">
                        <div class="d-flex align-items-center">
                          <span class="badge bg-primary me-2">条件组 {{ groupIndex + 1 }}</span>
                          <div class="btn-group btn-group-sm" role="group">
                            <input type="radio" class="btn-check" :id="`groupLogic${group.id}And`" v-model="group.logic" value="AND">
                            <label class="btn btn-outline-primary" :for="`groupLogic${group.id}And`">AND</label>
                            
                            <input type="radio" class="btn-check" :id="`groupLogic${group.id}Or`" v-model="group.logic" value="OR">
                            <label class="btn btn-outline-success" :for="`groupLogic${group.id}Or`">OR</label>
                          </div>
                        </div>
                        <div class="d-flex align-items-center gap-2">
                          <!-- 组间连接符（除了第一个组） -->
                          <div v-if="groupIndex > 0" class="d-flex align-items-center">
                            <span class="text-muted me-2">与上一组:</span>
                            <div class="btn-group btn-group-sm" role="group">
                              <input type="radio" class="btn-check" :id="`groupConnector${group.id}And`" v-model="group.connector" value="AND">
                              <label class="btn btn-outline-primary" :for="`groupConnector${group.id}And`">AND</label>
                              
                              <input type="radio" class="btn-check" :id="`groupConnector${group.id}Or`" v-model="group.connector" value="OR">
                              <label class="btn btn-outline-success" :for="`groupConnector${group.id}Or`">OR</label>
                            </div>
                          </div>
                          <button type="button" class="btn btn-outline-danger btn-sm" 
                                  @click="removePolicyGroup(groupIndex)"
                                  :disabled="policyGroups.length <= 1">
                            <i class="bi bi-trash"></i>
                          </button>
                        </div>
                      </div>

                      <!-- 组内条件列表 -->
                      <div v-if="group.conditions.length > 0">
                        <div v-for="(condition, condIndex) in group.conditions" :key="condition.id" 
                             class="policy-condition mb-2 p-2 border rounded"
                             :data-attribute="condition.attribute">
                          <div class="row align-items-center">
                            <div class="col-md-3">
                              <label class="form-label small mb-1">属性类型</label>
                              <select class="form-select form-select-sm" v-model="condition.attribute">
                                <option value="">选择属性</option>
                                <option value="wallet">钱包地址</option>
                                <option value="name">医生姓名</option>
                                <option value="department">科室</option>
                                <option value="hospital">医院</option>
                                <option value="title">职称</option>
                                <option value="license">执业编号</option>
                                <option value="specialties">专长</option>
                                <option value="did">DID标识符</option>
                              </select>
                            </div>
                            <div class="col-md-2">
                              <label class="form-label small mb-1">操作符</label>
                              <select class="form-select form-select-sm" v-model="condition.operator">
                                <option value="equals">等于</option>
                                <option value="contains">包含</option>
                                <option value="startsWith">开始于</option>
                                <option value="endsWith">结束于</option>
                              </select>
                            </div>
                            <div class="col-md-5">
                              <label class="form-label small mb-1">属性值</label>
                              <!-- 根据属性类型显示不同的输入控件 -->
                              <select v-if="isSelectableAttribute(condition.attribute)" 
                                      class="form-select form-select-sm" 
                                      v-model="condition.value">
                                <option value="">请选择{{ getAttributeLabel(condition.attribute) }}</option>
                                <option v-for="option in getAttributeOptions(condition.attribute)" 
                                        :key="option" 
                                        :value="option">
                                  {{ option }}
                                </option>
                              </select>
                              <input v-else 
                                     type="text" 
                                     class="form-control form-control-sm" 
                                     v-model="condition.value" 
                                     :placeholder="getAttributePlaceholder(condition.attribute)">
                            </div>
                            <div class="col-md-2">
                              <label class="form-label small mb-1">&nbsp;</label>
                              <div class="d-flex gap-1">
                                <button type="button" class="btn btn-outline-danger btn-sm" 
                                        @click="removeConditionFromGroup(groupIndex, condIndex)"
                                        :disabled="group.conditions.length <= 1">
                                  <i class="bi bi-trash"></i>
                                </button>
                                <button type="button" class="btn btn-outline-secondary btn-sm" 
                                        @click="duplicateConditionInGroup(groupIndex, condIndex)">
                                  <i class="bi bi-files"></i>
                                </button>
                              </div>
                            </div>
                          </div>
                          
                          <!-- 条件预览 -->
                          <div class="mt-2" v-if="condition.attribute && condition.value">
                            <small class="text-muted">
                              条件预览: <code>{{ formatConditionPreview(condition) }}</code>
                            </small>
                          </div>
                        </div>
                      </div>

                      <!-- 在组内添加条件 -->
                      <div class="mt-2">
                        <button type="button" class="btn btn-sm btn-outline-primary" @click="addConditionToGroup(groupIndex)">
                          <i class="bi bi-plus-circle me-1"></i>在此组添加条件
                        </button>
                      </div>
                    </div>
                  </div>

                  <!-- 策略预览 -->
                  <div class="mt-3 pt-3 border-top">
                    <div class="d-flex justify-content-between align-items-center">
                      <h6 class="mb-0">最终策略预览:</h6>
                    </div>
                    <div class="mt-2 p-3 bg-light border rounded">
                      <code class="text-dark">{{ complexPolicyPreview }}</code>
                    </div>
                  </div>

                  <!-- 快速模板 -->
                  <div class="mt-3 pt-3 border-top">
                    <h6 class="mb-2"><i class="bi bi-lightning me-1"></i>快速模板</h6>
                    <div class="btn-group btn-group-sm flex-wrap" role="group">
                      <button type="button" class="btn btn-outline-info" @click="applyDepartmentTemplate">
                        单科室限制
                      </button>
                      <button type="button" class="btn btn-outline-info" @click="applyHospitalTemplate">
                        单医院限制
                      </button>
                      <button type="button" class="btn btn-outline-info" @click="applySeniorDoctorTemplate">
                        高级医师(OR)
                      </button>
                      <button type="button" class="btn btn-outline-info" @click="applyDoctorAndDeptTemplate">
                        医院+科室组合
                      </button>
                      <button type="button" class="btn btn-outline-warning" @click="clearAllConditions">
                        重置
                      </button>
                    </div>
                    <div class="mt-2">
                      <small class="text-muted">
                        <i class="bi bi-info-circle me-1"></i>
                        "医院+科室组合"示例：北京协和医院 AND (心内科 OR 心外科)
                      </small>
                    </div>
                  </div>
                </div>

                <div class="form-text">
                  <i class="bi bi-info-circle me-1"></i>
                  设置访问控制策略，只有满足条件的用户才能访问此NFT元数据。支持基于VC属性的细粒度控制。
                </div>
              </div>

              <div class="mb-3">
                <label for="ciphertext" class="form-label">密文</label>
                <textarea class="form-control" id="ciphertext" v-model="newMetadata.ciphertext" rows="3"
                  required></textarea>
                <div class="form-text">加密后的内容（如果没有，请输入"none"）</div>
              </div>

              <div class="d-grid">
                <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                  <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                    aria-hidden="true"></span>
                  创建元数据
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            预览
          </div>
          <div class="card-body">
            <div v-if="newMetadata.name || newMetadata.description || newMetadata.image" class="metadata-preview">
              <img v-if="newMetadata.image" :src="getImageDisplayUrl(newMetadata.image)" class="img-fluid mb-3" alt="NFT Preview"
                @error="handleImageError">
              <img v-else
                src="data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E"
                class="img-fluid mb-3" alt="NFT Preview">

              <h4>{{ newMetadata.name || 'NFT 名称' }}</h4>
              <p>{{ newMetadata.description || '暂无描述' }}</p>

              <div>
                <h5 class="mt-3">访问控制</h5>
                <div class="row g-2">
                  <!-- 策略属性 -->
                  <div v-if="complexPolicyPreview && complexPolicyPreview !== '请添加访问条件'" class="col-12">
                    <div class="attribute-box p-2 border rounded bg-light">
                      <div class="text-muted small">访问策略</div>
                      <div class="fw-bold">{{ complexPolicyPreview }}</div>
                    </div>
                  </div>

                  <!-- 密文 -->
                  <div v-if="newMetadata.ciphertext" class="col-12 mt-2">
                    <div class="attribute-box p-2 border rounded bg-light">
                      <div class="text-muted small">密文</div>
                      <div class="fw-bold">{{ newMetadata.ciphertext.length > 20 ? newMetadata.ciphertext.substring(0,
                        20) + '...' : newMetadata.ciphertext }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="text-center p-5 text-muted">
              <i class="bi bi-card-image fs-1"></i>
              <p class="mt-3">填写表单以查看预览</p>
            </div>
          </div>
        </div>

        <div class="card mt-3" v-if="metadataUrl">
          <div class="card-header">
            元数据 URL
          </div>
          <div class="card-body">
            <div class="input-group">
              <input type="text" class="form-control" :value="metadataUrl" readonly>
              <button class="btn btn-outline-secondary" type="button" @click="copyToClipboard(metadataUrl)">
                <i class="bi bi-clipboard"></i>
              </button>
            </div>
            <div class="mt-3">
              <a :href="metadataUrl" target="_blank" class="btn btn-sm btn-outline-primary">
                <i class="bi bi-box-arrow-up-right me-1"></i>打开链接
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { useStore } from 'vuex'
import abeService from '@/services/abeService'

export default {
  name: 'MetadataManager',
  setup() {
    const store = useStore()

    const newMetadata = reactive({
        name: '',
        description: '',
        image: '',
        policy: '',
        ciphertext: 'none',
    })
    
    const isSubmitting = ref(false)
    const isUploadingImage = ref(false)
    const metadataUrl = ref('')

    // 策略编辑器相关数据
    const policyGroups = ref([])
    
    // 创建初始条件
    const createNewCondition = () => ({
      id: Date.now() + Math.random(),
      attribute: '',
      operator: 'equals',
      value: ''
    })

    // 创建新的策略组
    const createNewGroup = () => ({
      id: Date.now() + Math.random(),
      logic: 'AND',  // 组内逻辑
      connector: 'AND',  // 与上一组的连接符
      conditions: [createNewCondition()]
    })

    // 预定义选项数据
    const attributeOptions = {
      department: [
        '心内科', '心外科', '神经内科', '神经外科', '呼吸科', '消化科', 
        '肾内科', '内分泌科', '血液科', '肿瘤科', '儿科', '妇产科', 
        '骨科', '泌尿外科', '胸外科', '普外科', '整形外科', '眼科', 
        '耳鼻喉科', '口腔科', '皮肤科', '麻醉科', '影像科', '检验科', 
        '病理科', '药剂科', '康复科', '中医科', '急诊科', 'ICU重症监护'
      ],
      hospital: [
        '北京协和医院', '北京同仁医院', '北京安贞医院', '北京朝阳医院', 
        '北京友谊医院', '北京天坛医院', '北京宣武医院', '北京大学第一医院',
        '北京大学人民医院', '北京大学第三医院', '清华大学附属医院', '北京中医医院',
        '上海瑞金医院', '上海华山医院', '上海中山医院', '上海第一人民医院',
        '上海第六人民医院', '上海东方医院', '复旦大学附属医院', '上海交通大学医学院附属医院',
        '广州中山大学附属医院', '广州南方医院', '深圳人民医院', '深圳第二人民医院',
        '天津医科大学总医院', '天津第一中心医院', '西安交通大学第一附属医院',
        '四川大学华西医院', '重庆医科大学附属医院', '山东大学齐鲁医院'
      ],
      title: [
        '住院医师', '主治医师', '副主任医师', '主任医师',
        '高级医师', '特级专家', '首席专家', '学科带头人',
        '实习医师', '规培医师', '进修医师', '访问学者'
      ],
      specialties: [
        '心血管疾病', '脑血管疾病', '呼吸系统疾病', '消化系统疾病',
        '内分泌疾病', '肿瘤治疗', '外科手术', '微创手术', 
        '介入治疗', '康复治疗', '儿童疾病', '妇科疾病',
        '骨科疾病', '皮肤病', '眼科疾病', '耳鼻喉疾病',
        '口腔疾病', '精神疾病', '中医治疗', '针灸推拿',
        '急救医学', '重症医学', '麻醉医学', '影像诊断',
        '病理诊断', '检验医学', '药物治疗', '预防医学'
      ]
    }

    // 判断是否为可选择的属性
    const isSelectableAttribute = (attribute) => {
      return ['department', 'hospital', 'title', 'specialties'].includes(attribute)
    }

    // 获取属性选项
    const getAttributeOptions = (attribute) => {
      return attributeOptions[attribute] || []
    }

    // 获取属性标签
    const getAttributeLabel = (attribute) => {
      const labels = {
        department: '科室',
        hospital: '医院',
        title: '职称',
        specialties: '专长'
      }
      return labels[attribute] || '值'
    }

    // 获取属性占位符
    const getAttributePlaceholder = (attribute) => {
      const placeholders = {
        wallet: '0x...',
        name: '张医生',
        department: '选择科室',
        hospital: '选择医院',
        title: '选择职称',
        license: '110101199001011234',
        specialties: '选择专长',
        did: 'did:ethr:0x...'
      }
      return placeholders[attribute] || '请输入值'
    }

    // 初始化默认策略组
    const initializePolicyGroups = () => {
      if (policyGroups.value.length === 0) {
        policyGroups.value.push(createNewGroup())
      }
    }

    // 添加策略组
    const addPolicyGroup = () => {
      policyGroups.value.push(createNewGroup())
    }

    // 移除策略组
    const removePolicyGroup = (groupIndex) => {
      if (policyGroups.value.length > 1) {
        policyGroups.value.splice(groupIndex, 1)
      }
    }

    // 在指定组添加条件
    const addConditionToGroup = (groupIndex) => {
      policyGroups.value[groupIndex].conditions.push(createNewCondition())
    }

    // 从组中移除条件
    const removeConditionFromGroup = (groupIndex, condIndex) => {
      const group = policyGroups.value[groupIndex]
      if (group.conditions.length > 1) {
        group.conditions.splice(condIndex, 1)
      }
    }

    // 在组中复制条件
    const duplicateConditionInGroup = (groupIndex, condIndex) => {
      const group = policyGroups.value[groupIndex]
      const original = group.conditions[condIndex]
      const copy = {
        ...original,
        id: Date.now() + Math.random()
      }
      group.conditions.splice(condIndex + 1, 0, copy)
    }

    // 格式化条件预览
    const formatConditionPreview = (condition) => {
      if (!condition.attribute || !condition.value) return ''
      
      const operators = {
        equals: '==',
        contains: 'contains',
        startsWith: 'startsWith',
        endsWith: 'endsWith'
      }
      
      return `${condition.attribute} ${operators[condition.operator]} "${condition.value}"`
    }

    // 复杂策略预览
    const complexPolicyPreview = computed(() => {
      const groupStrings = []
      
      for (let i = 0; i < policyGroups.value.length; i++) {
        const group = policyGroups.value[i]
        const validConditions = group.conditions.filter(c => c.attribute && c.value)
        
        if (validConditions.length === 0) continue
        
        const conditionStrings = validConditions.map(formatConditionPreview)
        let groupString = ''
        
        if (conditionStrings.length === 1) {
          groupString = conditionStrings[0]
        } else {
          groupString = `(${conditionStrings.join(` ${group.logic} `)})`
        }
        
        // 添加组间连接符
        if (i > 0 && groupStrings.length > 0) {
          groupStrings.push(` ${group.connector} `)
        }
        
        groupStrings.push(groupString)
      }
      
      return groupStrings.length > 0 ? groupStrings.join('') : '请添加访问条件'
    })

    // 生成实际的ABE策略字符串
    const generateABEPolicy = () => {
      const groupStrings = []
      
      for (let i = 0; i < policyGroups.value.length; i++) {
        const group = policyGroups.value[i]
        const validConditions = group.conditions.filter(c => c.attribute && c.value)
        
        if (validConditions.length === 0) continue
        
        // 转换为ABE策略格式
        const abeConditions = validConditions.map(condition => {
          return `${condition.attribute}:${condition.value}`
        })
        
        let groupString = ''
        if (abeConditions.length === 1) {
          groupString = abeConditions[0]
        } else {
          groupString = `(${abeConditions.join(` ${group.logic} `)})`
        }
        
        // 添加组间连接符
        if (i > 0 && groupStrings.length > 0) {
          groupStrings.push(` ${group.connector} `)
        }
        
        groupStrings.push(groupString)
      }
      
      if (groupStrings.length === 0) {
        return 'department:心内科'  // 默认策略
      }
      
      const result = groupStrings.join('')
      
      // 如果有多个组，外层添加括号
      if (policyGroups.value.filter(g => g.conditions.filter(c => c.attribute && c.value).length > 0).length > 1) {
        return `(${result})`
      }
      
      return result
    }

    // 快速模板方法
    const applyDepartmentTemplate = () => {
      policyGroups.value = [
        {
          id: Date.now(),
          logic: 'OR',
          connector: 'AND',
          conditions: [
            {
              id: Date.now(),
              attribute: 'department',
              operator: 'equals',
              value: '心内科'
            }
          ]
        }
      ]
    }

    const applyHospitalTemplate = () => {
      policyGroups.value = [
        {
          id: Date.now(),
          logic: 'AND',
          connector: 'AND',
          conditions: [
            {
              id: Date.now(),
              attribute: 'hospital',
              operator: 'equals',
              value: '北京协和医院'
            }
          ]
        }
      ]
    }

    const applySeniorDoctorTemplate = () => {
      policyGroups.value = [
        {
          id: Date.now(),
          logic: 'OR',
          connector: 'AND',
          conditions: [
            {
              id: Date.now(),
              attribute: 'title',
              operator: 'equals',
              value: '主任医师'
            },
            {
              id: Date.now() + 1,
              attribute: 'title',
              operator: 'equals',
              value: '副主任医师'
            }
          ]
        }
      ]
    }

    const applyDoctorAndDeptTemplate = () => {
      // 医院A AND (科室B OR 科室C) 的示例
      policyGroups.value = [
        {
          id: Date.now(),
          logic: 'AND',
          connector: 'AND',
          conditions: [
            {
              id: Date.now(),
              attribute: 'hospital',
              operator: 'equals',
              value: '北京协和医院'
            }
          ]
        },
        {
          id: Date.now() + 1,
          logic: 'OR',
          connector: 'AND',
          conditions: [
            {
              id: Date.now() + 2,
              attribute: 'department',
              operator: 'equals',
              value: '心内科'
            },
            {
              id: Date.now() + 3,
              attribute: 'department',
              operator: 'equals',
              value: '心外科'
            }
          ]
        }
      ]
    }

    const clearAllConditions = () => {
      policyGroups.value = [createNewGroup()]
    }

    // 监听策略变化，自动更新newMetadata.policy
    watch(policyGroups, () => {
      newMetadata.policy = generateABEPolicy()
    }, { deep: true })

    // 处理图片加载错误
    const handleImageError = (event) => {
      if (!event.target.src.includes('placeholder')) {
        event.target.src = 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E';
        event.target.onerror = null;
      }
    }

    // 处理图片选择
    const handleImageSelect = async (event) => {
      const file = event.target.files[0]
      if (!file) return

      isUploadingImage.value = true
      try {
        console.log('开始上传图片:', file.name)
        const result = await abeService.uploadImage(file)
        
        // 使用主要的HTTP URL作为图片地址
        newMetadata.image = result.primary_url
        
        store.commit('notifications/add', {
          type: 'success',
          message: `图片上传成功！IPFS Hash: ${result.hash}`
        })
        console.log('图片上传结果:', result)
      } catch (error) {
        console.error('图片上传失败:', error)
        store.commit('notifications/add', {
          type: 'danger',
          message: '图片上传失败: ' + error.message
        })
      } finally {
        isUploadingImage.value = false
        // 清空input的value，允许重新选择同一个文件
        event.target.value = ''
      }
    }

    // 移除图片
    const removeImage = () => {
      newMetadata.image = ''
      store.commit('notifications/add', {
        type: 'success',
        message: '图片已移除'
      })
    }

    // 获取图片显示URL
    const getImageDisplayUrl = (imageUrl) => {
      if (!imageUrl) return ''
      
      // 如果是IPFS链接，转换为HTTP网关链接
      if (imageUrl.startsWith('ipfs://')) {
        const hash = imageUrl.replace('ipfs://', '')
        return `https://dweb.link/ipfs/${hash}`
      }
      
      return imageUrl
    }

    const createMetadata = async () => {
      // 先生成最新的策略
      newMetadata.policy = generateABEPolicy()
      
      if (!newMetadata.name || !newMetadata.description || !newMetadata.image || !newMetadata.policy || !newMetadata.ciphertext) {
        store.commit('notifications/add', {
          type: 'danger',
          message: '请填写所有必填字段'
        });
        return;
      }

      // 检查是否有有效的策略条件
      const hasValidConditions = policyGroups.value.some(group => 
        group.conditions.some(c => c.attribute && c.value)
      )
      if (!hasValidConditions) {
        store.commit('notifications/add', {
          type: 'danger',
          message: '请至少添加一个有效的访问条件'
        });
        return;
      }

      isSubmitting.value = true;

      try {
        const metadata = {
          name: newMetadata.name,
          description: newMetadata.description,
          image: newMetadata.image,
          external_url: '',
          policy: newMetadata.policy,
          ciphertext: newMetadata.ciphertext
        };

        console.log('提交的元数据:', metadata);

        // 调用API上传元数据
        const response = await fetch(`${process.env.VUE_APP_API_BASE_URL || 'http://localhost:8080'}/api/metadata`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(metadata)
        });

        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.error || '上传元数据失败');
        }

        const data = await response.json();
        metadataUrl.value = `ipfs://${data.ipfs_hash}`;

        store.commit('notifications/add', {
          type: 'success',
          message: '元数据创建成功! IPFS哈希: ' + data.ipfs_hash
        });
      } catch (error) {
        console.error('创建元数据失败:', error);
        store.commit('notifications/add', {
          type: 'danger',
          message: '创建元数据失败: ' + error.message
        });
      } finally {
        isSubmitting.value = false;
      }
    }

    const copyToClipboard = (text) => {
      navigator.clipboard.writeText(text).then(
        () => {
          store.commit('notifications/add', {
            type: 'success',
            message: 'URL已复制到剪贴板'
          });
        },
        () => {
          store.commit('notifications/add', {
            type: 'danger',
            message: '复制到剪贴板失败'
          });
        }
      );
    }

    // 组件挂载时初始化
    onMounted(() => {
      initializePolicyGroups()
      // 初始化默认策略
      newMetadata.policy = generateABEPolicy()
    })

    return {
      newMetadata,
      isSubmitting,
      isUploadingImage,
      metadataUrl,
      policyGroups,
      complexPolicyPreview,
      isSelectableAttribute,
      getAttributeOptions,
      getAttributeLabel,
      getAttributePlaceholder,
      addPolicyGroup,
      removePolicyGroup,
      addConditionToGroup,
      removeConditionFromGroup,
      duplicateConditionInGroup,
      formatConditionPreview,
      applyDepartmentTemplate,
      applyHospitalTemplate,
      applySeniorDoctorTemplate,
      applyDoctorAndDeptTemplate,
      clearAllConditions,
      handleImageError,
      handleImageSelect,
      removeImage,
      getImageDisplayUrl,
      createMetadata,
      copyToClipboard
    }
  }
};
</script>

<style scoped>
.attribute-box {
  background-color: rgba(0, 0, 0, 0.03);
  transition: all 0.2s;
}

.attribute-box:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.metadata-preview img {
  max-height: 300px;
  object-fit: contain;
  width: 100%;
}

.image-upload-section {
  border: 1px dashed #dee2e6;
  border-radius: 0.5rem;
  padding: 1rem;
  background-color: #f8f9fa;
}

.image-preview-container {
  position: relative;
  display: inline-block;
  max-width: 300px;
}

.image-preview {
  max-width: 100%;
  max-height: 200px;
  border-radius: 0.5rem;
  border: 1px solid #dee2e6;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

/* 策略编辑器样式 */
.policy-editor {
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
  border: 2px solid #dee2e6 !important;
}

/* 策略组样式 */
.policy-group {
  background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
  border: 2px solid #e9ecef !important;
  position: relative;
}

.policy-group::before {
  content: '';
  position: absolute;
  left: -3px;
  top: 0;
  bottom: 0;
  width: 4px;
  background: linear-gradient(180deg, #28a745 0%, #20c997 100%);
  border-radius: 0 3px 3px 0;
}

.policy-group:hover {
  border-color: #28a745 !important;
  box-shadow: 0 4px 12px rgba(40, 167, 69, 0.15);
}

/* 组标题徽章 */
.policy-group .badge {
  font-size: 0.875rem;
  padding: 0.5rem 0.75rem;
}

.policy-condition {
  background: #ffffff;
  border: 1px solid #dee2e6;
  transition: all 0.2s ease;
  position: relative;
}

.policy-condition:hover {
  border-color: #0d6efd;
  box-shadow: 0 0 0 0.2rem rgba(13, 110, 253, 0.1);
}

.policy-condition::before {
  content: '';
  position: absolute;
  left: -3px;
  top: 0;
  bottom: 0;
  width: 3px;
  background: linear-gradient(180deg, #0d6efd 0%, #6f42c1 100%);
  border-radius: 0 2px 2px 0;
}

.btn-group-sm .btn {
  font-size: 0.775rem;
}

/* 条件预览样式 */
.policy-condition code {
  background: #f8f9fa;
  color: #495057;
  padding: 0.2rem 0.4rem;
  border-radius: 0.25rem;
  font-size: 0.875rem;
}

/* 快速模板按钮样式 */
.btn-group .btn-outline-info {
  border-color: #0dcaf0;
  color: #0dcaf0;
}

.btn-group .btn-outline-info:hover {
  background-color: #0dcaf0;
  border-color: #0dcaf0;
  color: #fff;
}

.btn-group .btn-outline-warning {
  border-color: #ffc107;
  color: #ffc107;
}

.btn-group .btn-outline-warning:hover {
  background-color: #ffc107;
  border-color: #ffc107;
  color: #000;
}

/* 不同属性类型的颜色区分 */
.policy-condition[data-attribute="department"] {
  border-left: 3px solid #20c997;
}

.policy-condition[data-attribute="hospital"] {
  border-left: 3px solid #fd7e14;
}

.policy-condition[data-attribute="title"] {
  border-left: 3px solid #6f42c1;
}

.policy-condition[data-attribute="specialties"] {
  border-left: 3px solid #e83e8c;
}

/* 属性标签样式 */
.policy-condition .form-label {
  font-weight: 600;
  color: #495057;
}

/* 选择框动画 */
.policy-condition .form-select {
  transition: all 0.3s ease;
}

.policy-condition .form-select:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* 策略预览高亮 */
.text-muted code {
  background: linear-gradient(90deg, #e3f2fd 0%, #f3e5f5 100%);
  border: 1px solid #e1bee7;
  padding: 0.3rem 0.6rem;
  border-radius: 0.4rem;
  font-weight: 500;
}

/* 动画效果 */
.policy-condition {
  animation: slideIn 0.3s ease-out;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .policy-condition .row > div {
    margin-bottom: 0.5rem;
  }
  
  .btn-group {
    flex-wrap: wrap;
  }
  
  .btn-group .btn {
    margin: 0.1rem;
  }
}
</style>