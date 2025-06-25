<template>
  <div class="doctor-did">
    <h2><i class="bi bi-hospital me-2"></i>医生DID集成</h2>

    <div class="alert alert-info">
      <i class="bi bi-info-circle me-2"></i>
      <strong>医生DID集成</strong>允许医疗机构为医生创建分布式身份标识符(DID)和可验证凭证(VC)，以证明其专业资格和执业权限。
    </div>

    <div class="row">
      <div class="col-md-6">
        <div class="card mb-4">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-person-badge me-2"></i>医生DID创建</h5>
          </div>
          <div class="card-body">
            <form @submit.prevent="createDoctorDID">
              <div class="mb-3">
                <label for="doctor-name" class="form-label">医生姓名</label>
                <input type="text" class="form-control" id="doctor-name" v-model="doctorForm.name" placeholder="输入医生姓名"
                  required>
              </div>

              <div class="mb-3">
                <label for="doctor-id" class="form-label">医生ID</label>
                <input type="text" class="form-control" id="doctor-id" v-model="doctorForm.id" placeholder="输入医生ID"
                  required>
                <div class="form-text">
                  医生的唯一标识符，如医师执业证号。
                </div>
              </div>

              <div class="mb-3">
                <label for="doctor-department" class="form-label">科室</label>
                <select class="form-select" id="doctor-department" v-model="doctorForm.department" required>
                  <option value="" disabled>选择科室</option>
                  <option v-for="dept in departmentOptions" :key="dept" :value="dept">{{ dept }}</option>
                </select>
              </div>

              <div class="mb-3">
                <label for="doctor-hospital" class="form-label">医院</label>
                <select class="form-select" id="doctor-hospital" v-model="doctorForm.hospital" required>
                  <option value="" disabled>选择医院</option>
                  <option v-for="hospital in hospitalOptions" :key="hospital" :value="hospital">{{ hospital }}</option>
                </select>
              </div>

              <div class="mb-3">
                <label for="doctor-title" class="form-label">职称</label>
                <select class="form-select" id="doctor-title" v-model="doctorForm.title" required>
                  <option value="" disabled>选择职称</option>
                  <option v-for="title in titleOptions" :key="title" :value="title">{{ title }}</option>
                </select>
              </div>

              <div class="mb-3">
                <label for="doctor-license" class="form-label">执业证书编号</label>
                <input type="text" class="form-control" id="doctor-license" v-model="doctorForm.license"
                  placeholder="输入执业证书编号" required>
              </div>

              <div class="mb-3">
                <label for="doctor-specialties" class="form-label">专长</label>
                <select class="form-select" id="doctor-specialties" v-model="doctorForm.specialties" multiple>
                  <option v-for="specialty in specialtyOptions" :key="specialty" :value="specialty">{{ specialty }}</option>
                </select>
                <div class="form-text">
                  可多选。按住Ctrl/Cmd键点击可选择多个专长。
                </div>
              </div>

              <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                  aria-hidden="true"></span>
                <i v-else class="bi bi-person-plus me-1"></i>创建医生DID
              </button>
            </form>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card mb-4" v-if="createdDoctorDID">
          <div class="card-header bg-success text-white">
            <h5 class="mb-0"><i class="bi bi-check-circle me-2"></i>医生DID创建成功</h5>
          </div>
          <div class="card-body">
            <div class="mb-3">
              <label class="form-label">医生DID</label>
              <div class="input-group">
                <input type="text" class="form-control" readonly :value="createdDoctorDID.did">
                <button class="btn btn-outline-primary" type="button" @click="copyToClipboard(createdDoctorDID.did)">
                  <i class="bi bi-clipboard"></i>
                </button>
              </div>
            </div>

            <div class="mb-3">
              <h6>医生信息</h6>
              <div class="table-responsive">
                <table class="table table-bordered">
                  <tbody>
                    <tr>
                      <th>姓名</th>
                      <td>{{ doctorForm.name }}</td>
                    </tr>
                    <tr>
                      <th>医生ID</th>
                      <td>{{ doctorForm.id }}</td>
                    </tr>
                    <tr>
                      <th>科室</th>
                      <td>{{ doctorForm.department }}</td>
                    </tr>
                    <tr>
                      <th>医院</th>
                      <td>{{ doctorForm.hospital }}</td>
                    </tr>
                    <tr>
                      <th>职称</th>
                      <td>{{ doctorForm.title }}</td>
                    </tr>
                    <tr>
                      <th>执业证书编号</th>
                      <td>{{ doctorForm.license }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div class="d-grid gap-2">
              <button class="btn btn-primary" @click="issueCredential" :disabled="isIssuingVC">
                <span v-if="isIssuingVC" class="spinner-border spinner-border-sm me-1" role="status"
                  aria-hidden="true"></span>
                <i v-else class="bi bi-card-checklist me-1"></i>颁发医生凭证
              </button>
            </div>
          </div>
        </div>

        <div class="card mb-4" v-if="createdVC">
          <div class="card-header bg-success text-white">
            <h5 class="mb-0"><i class="bi bi-check-circle me-2"></i>医生凭证颁发成功</h5>
          </div>
          <div class="card-body">
            <div class="mb-3">
              <label class="form-label">凭证ID</label>
              <div class="input-group">
                <input type="text" class="form-control" readonly :value="createdVC.id || createdVC.vcId || '未生成'">
                <button class="btn btn-outline-primary" type="button" @click="copyToClipboard(createdVC.id || createdVC.vcId)">
                  <i class="bi bi-clipboard"></i>
                </button>
              </div>
            </div>

            <div class="mb-3">
              <label class="form-label">凭证内容</label>
              <pre class="bg-light p-3 rounded">{{ formattedVC }}</pre>
            </div>

            <div class="d-grid gap-2">
              <button class="btn btn-primary" @click="downloadVC">
                <i class="bi bi-download me-1"></i>下载凭证
              </button>
              <router-link :to="{ name: 'DIDList' }" class="btn btn-outline-primary">
                <i class="bi bi-card-checklist me-1"></i>查看我的凭证
              </router-link>
            </div>
          </div>
        </div>

        <div class="card" v-if="!createdDoctorDID">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-info-circle me-2"></i>医生DID集成说明</h5>
          </div>
          <div class="card-body">
            <h6>什么是医生DID?</h6>
            <p>
              医生DID是为医疗专业人员创建的分布式身份标识符，用于证明其身份和专业资格。通过DID和可验证凭证，医生可以安全地证明其专业资格，而无需依赖中央机构。
            </p>

            <h6>使用场景</h6>
            <ul>
              <li><strong>医疗机构认证:</strong> 医院可以为其医生颁发可验证凭证，证明其在该医院的执业资格。</li>
              <li><strong>跨机构协作:</strong> 医生可以使用DID在不同医疗机构之间安全地共享其专业资格。</li>
              <li><strong>远程医疗:</strong> 患者可以验证远程医生的身份和资格。</li>
              <li><strong>医学研究:</strong> 研究机构可以验证参与研究的医生的专业背景。</li>
            </ul>

            <h6>创建流程</h6>
            <ol>
              <li>填写医生基本信息</li>
              <li>系统为医生创建DID</li>
              <li>医疗机构作为颁发者为医生颁发可验证凭证</li>
              <li>医生可以使用DID和凭证证明其身份和专业资格</li>
            </ol>
          </div>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="col-md-12">
        <div class="card">
          <div class="card-header d-flex justify-content-between align-items-center">
            <h5 class="mb-0"><i class="bi bi-list-check me-2"></i>医生DID列表</h5>
            <button class="btn btn-sm btn-outline-primary" @click="loadDoctorDIDs">
              <i class="bi bi-arrow-clockwise me-1"></i>刷新
            </button>
          </div>
          <div class="card-body">
            <!-- 加载状态 -->
            <div v-if="isLoadingDoctors" class="text-center my-3">
              <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">加载中...</span>
              </div>
              <p class="mt-2">加载医生DID列表...</p>
            </div>

            <!-- 没有医生DID时显示 -->
            <div v-else-if="!doctorDIDs.length" class="text-center my-3">
              <i class="bi bi-person-badge fs-1 text-muted"></i>
              <p class="lead mt-3">暂无医生DID记录</p>
            </div>

            <!-- 医生DID列表 -->
            <div v-else class="table-responsive">
              <table class="table">
                <thead>
                  <tr>
                    <th>姓名</th>
                    <th>医生ID</th>
                    <th>科室</th>
                    <th>医院</th>
                    <th>DID</th>
                    <th>凭证状态</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="doctor in doctorDIDs" :key="doctor.id">
                    <td>{{ doctor.name }}</td>
                    <td>{{ doctor.doctorId }}</td>
                    <td>{{ doctor.department }}</td>
                    <td>{{ doctor.hospital }}</td>
                    <td>
                      <span :title="doctor.did" class="did-text">{{ truncateDID(doctor.did) }}</span>
                      <button class="btn btn-sm btn-outline-secondary ms-1" @click="copyToClipboard(doctor.did)">
                        <i class="bi bi-clipboard"></i>
                      </button>
                    </td>
                    <td>
                      <span :class="['badge', doctor.hasVC ? 'bg-success' : 'bg-warning']">
                        {{ doctor.hasVC ? '已颁发' : '未颁发' }}
                      </span>
                    </td>
                    <td>
                      <button class="btn btn-sm btn-primary me-1" @click="viewDoctorDID(doctor)">
                        <i class="bi bi-eye"></i>
                      </button>
                      <button v-if="!doctor.hasVC" class="btn btn-sm btn-success"
                        @click="issueDoctorCredential(doctor)">
                        <i class="bi bi-card-checklist"></i>
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 医生DID详情模态框 -->
    <div class="modal fade" ref="doctorDetailModal" tabindex="-1">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">医生DID详情</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body" v-if="selectedDoctor">
            <div class="row">
              <div class="col-md-6">
                <h6>基本信息</h6>
                <div class="table-responsive">
                  <table class="table table-bordered">
                    <tbody>
                      <tr>
                        <th>姓名</th>
                        <td>{{ selectedDoctor.name }}</td>
                      </tr>
                      <tr>
                        <th>医生ID</th>
                        <td>{{ selectedDoctor.doctorId }}</td>
                      </tr>
                      <tr>
                        <th>科室</th>
                        <td>{{ selectedDoctor.department }}</td>
                      </tr>
                      <tr>
                        <th>医院</th>
                        <td>{{ selectedDoctor.hospital }}</td>
                      </tr>
                      <tr>
                        <th>职称</th>
                        <td>{{ selectedDoctor.title }}</td>
                      </tr>
                      <tr>
                        <th>执业证书编号</th>
                        <td>{{ selectedDoctor.license }}</td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="col-md-6">
                <h6>DID信息</h6>
                <p><strong>DID:</strong> {{ selectedDoctor.did }}</p>
                <p><strong>创建时间:</strong> {{ formatDate(selectedDoctor.createdAt) }}</p>
                <p>
                  <strong>凭证状态:</strong>
                  <span :class="['badge', selectedDoctor.hasVC ? 'bg-success' : 'bg-warning']">
                    {{ selectedDoctor.hasVC ? '已颁发' : '未颁发' }}
                  </span>
                </p>
                <p v-if="selectedDoctor.vcId"><strong>凭证ID:</strong> {{ selectedDoctor.vcId }}</p>
              </div>
            </div>

            <div class="mb-3" v-if="selectedDoctor.specialties">
              <h6>专长</h6>
              <p>{{ selectedDoctor.specialties }}</p>
            </div>

            <div class="mb-3" v-if="selectedDoctor.hasVC && selectedDoctor.vc">
              <h6>凭证内容</h6>
              <pre class="bg-light p-3 rounded">{{ JSON.stringify(selectedDoctor.vc, null, 2) }}</pre>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
            <button v-if="selectedDoctor && !selectedDoctor.hasVC" type="button" class="btn btn-success"
              @click="issueDoctorCredential(selectedDoctor)">
              <i class="bi bi-card-checklist me-1"></i>颁发凭证
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { Modal } from 'bootstrap'

export default {
  name: 'DoctorDID',

  setup() {
    const store = useStore()
    const doctorDetailModal = ref(null)
    let bsDoctorDetailModal = null

    // 预定义选项数据（与ABE加密页面保持一致）
    const departmentOptions = [
      '心内科', '心外科', '神经内科', '神经外科', '呼吸科', '消化科', 
      '肾内科', '内分泌科', '血液科', '肿瘤科', '儿科', '妇产科', 
      '骨科', '泌尿外科', '胸外科', '普外科', '整形外科', '眼科', 
      '耳鼻喉科', '口腔科', '皮肤科', '麻醉科', '影像科', '检验科', 
      '病理科', '药剂科', '康复科', '中医科', '急诊科', 'ICU重症监护'
    ]

    const hospitalOptions = [
      '北京协和医院', '北京同仁医院', '北京安贞医院', '北京朝阳医院', 
      '北京友谊医院', '北京天坛医院', '北京宣武医院', '北京大学第一医院',
      '北京大学人民医院', '北京大学第三医院', '清华大学附属医院', '北京中医医院',
      '上海瑞金医院', '上海华山医院', '上海中山医院', '上海第一人民医院',
      '上海第六人民医院', '上海东方医院', '复旦大学附属医院', '上海交通大学医学院附属医院',
      '广州中山大学附属医院', '广州南方医院', '深圳人民医院', '深圳第二人民医院',
      '天津医科大学总医院', '天津第一中心医院', '西安交通大学第一附属医院',
      '四川大学华西医院', '重庆医科大学附属医院', '山东大学齐鲁医院'
    ]

    const titleOptions = [
      '住院医师', '主治医师', '副主任医师', '主任医师',
      '高级医师', '特级专家', '首席专家', '学科带头人',
      '实习医师', '规培医师', '进修医师', '访问学者'
    ]

    const specialtyOptions = [
      '心血管疾病', '脑血管疾病', '呼吸系统疾病', '消化系统疾病',
      '内分泌疾病', '肿瘤治疗', '外科手术', '微创手术', 
      '介入治疗', '康复治疗', '儿童疾病', '妇科疾病',
      '骨科疾病', '皮肤病', '眼科疾病', '耳鼻喉疾病',
      '口腔疾病', '精神疾病', '中医治疗', '针灸推拿',
      '急救医学', '重症医学', '麻醉医学', '影像诊断',
      '病理诊断', '检验医学', '药物治疗', '预防医学'
    ]

    // 表单数据
    const doctorForm = ref({
      name: '',
      id: '',
      department: '',
      hospital: '',
      title: '',
      license: '',
      specialties: []
    })

    // 创建结果
    const createdDoctorDID = ref(null)
    const createdVC = ref(null)

    // 提交状态
    const isSubmitting = ref(false)
    const isIssuingVC = ref(false)

    // 医生DID列表
    const doctorDIDs = ref([])
    const isLoadingDoctors = ref(false)
    const selectedDoctor = ref(null)

    // 格式化的VC
    const formattedVC = computed(() => {
      if (!createdVC.value) return '{}'
      return JSON.stringify(createdVC.value, null, 2)
    })

    // 创建医生DID
    const createDoctorDID = async () => {
      if (!doctorForm.value.name || !doctorForm.value.id ||
        !doctorForm.value.department || !doctorForm.value.hospital ||
        !doctorForm.value.title || !doctorForm.value.license) {
        store.dispatch('app/showError', '请填写所有必填字段')
        return
      }

      isSubmitting.value = true

      try {
        const result = await store.dispatch('did/createDoctorDID', {
          name: doctorForm.value.name,
          doctorId: doctorForm.value.id,
          department: doctorForm.value.department,
          hospital: doctorForm.value.hospital,
          title: doctorForm.value.title,
          license: doctorForm.value.license,
          specialties: Array.isArray(doctorForm.value.specialties) 
            ? doctorForm.value.specialties.join(', ') 
            : doctorForm.value.specialties
        })

        if (result) {
          createdDoctorDID.value = result
          store.dispatch('app/showSuccess', '医生DID创建成功')

          // 重新加载医生DID列表
          loadDoctorDIDs()
        }
      } finally {
        isSubmitting.value = false
      }
    }

    // 颁发凭证
    const issueCredential = async () => {
      if (!createdDoctorDID.value) return

      isIssuingVC.value = true

      try {
        const result = await store.dispatch('did/issueDoctorCredential', {
          did: createdDoctorDID.value.did,
          name: doctorForm.value.name,
          doctorId: doctorForm.value.id,
          department: doctorForm.value.department,
          hospital: doctorForm.value.hospital,
          title: doctorForm.value.title,
          license: doctorForm.value.license,
          specialties: Array.isArray(doctorForm.value.specialties) 
            ? doctorForm.value.specialties.join(', ') 
            : doctorForm.value.specialties
        })

        if (result) {
          createdVC.value = result
          store.dispatch('app/showSuccess', '医生凭证颁发成功')

          // 重新加载医生DID列表
          loadDoctorDIDs()
        }
      } finally {
        isIssuingVC.value = false
      }
    }

    // 为选中的医生颁发凭证
    const issueDoctorCredential = async (doctor) => {
      if (!doctor) return

      // 如果当前在详情模态框中，先关闭它
      if (bsDoctorDetailModal) {
        bsDoctorDetailModal.hide()
      }

      isIssuingVC.value = true

      try {
        const result = await store.dispatch('did/issueDoctorCredential', {
          did: doctor.did,
          name: doctor.name,
          doctorId: doctor.doctorId,
          department: doctor.department,
          hospital: doctor.hospital,
          title: doctor.title,
          license: doctor.license,
          specialties: doctor.specialties // 这里不需要转换，因为从数据库来的已经是字符串
        })

        if (result) {
          createdVC.value = result
          store.dispatch('app/showSuccess', '医生凭证颁发成功')

          // 重新加载医生DID列表
          loadDoctorDIDs()
        }
      } finally {
        isIssuingVC.value = false
      }
    }

    // 加载医生DID列表
    const loadDoctorDIDs = async () => {
      isLoadingDoctors.value = true

      try {
        const result = await store.dispatch('did/loadDoctorDIDs')
        if (result) {
          doctorDIDs.value = result
        }
      } finally {
        isLoadingDoctors.value = false
      }
    }

    // 查看医生DID详情
    const viewDoctorDID = async (doctor) => {
      selectedDoctor.value = doctor

      // 如果医生有凭证但没有加载凭证内容，则加载凭证
      if (doctor.hasVC && doctor.vcId && !doctor.vc) {
        try {
          const vc = await store.dispatch('did/getVC', doctor.vcId)
          if (vc) {
            selectedDoctor.value = {
              ...selectedDoctor.value,
              vc
            }
          }
        } catch (error) {
          console.error('加载凭证失败:', error)
        }
      }

      if (bsDoctorDetailModal) {
        bsDoctorDetailModal.show()
      }
    }

    // 下载VC
    const downloadVC = () => {
      if (!createdVC.value) return

      const vcStr = JSON.stringify(createdVC.value, null, 2)
      const blob = new Blob([vcStr], { type: 'application/json' })
      const url = URL.createObjectURL(blob)

      const a = document.createElement('a')
      a.href = url
      // 确保能正确访问VC ID
      const vcId = createdVC.value.id || createdVC.value.vcId || 'unknown'
      const fileName = vcId.includes(':') ? vcId.split(':').pop() : vcId
      a.download = `doctor-vc-${fileName}.json`
      document.body.appendChild(a)
      a.click()
      document.body.removeChild(a)
      URL.revokeObjectURL(url)
    }

    // 复制到剪贴板
    const copyToClipboard = (text) => {
      navigator.clipboard.writeText(text).then(() => {
        store.dispatch('app/showSuccess', '已复制到剪贴板')
      }).catch(err => {
        console.error('复制失败:', err)
        store.dispatch('app/showError', '复制失败')
      })
    }

    // 截断DID
    const truncateDID = (did) => {
      if (!did) return ''
      if (did.length <= 30) return did
      return did.substring(0, 15) + '...' + did.substring(did.length - 10)
    }

    // 格式化日期
    const formatDate = (dateString) => {
      if (!dateString) return ''
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    onMounted(() => {
      // 加载医生DID列表
      loadDoctorDIDs()

      // 初始化模态框
      if (doctorDetailModal.value) {
        bsDoctorDetailModal = new Modal(doctorDetailModal.value)
      }
    })

    return {
      doctorForm,
      createdDoctorDID,
      createdVC,
      isSubmitting,
      isIssuingVC,
      doctorDIDs,
      isLoadingDoctors,
      selectedDoctor,
      formattedVC,
      doctorDetailModal,
      departmentOptions,
      hospitalOptions,
      titleOptions,
      specialtyOptions,
      createDoctorDID,
      issueCredential,
      issueDoctorCredential,
      loadDoctorDIDs,
      viewDoctorDID,
      downloadVC,
      copyToClipboard,
      truncateDID,
      formatDate
    }
  }
}
</script>

<style scoped>
.did-text {
  font-family: monospace;
}

pre {
  max-height: 300px;
  overflow-y: auto;
}

.card {
  margin-bottom: 1.5rem;
}

/* 多选下拉框样式 */
select[multiple] {
  min-height: 120px;
  background-image: none;
}

select[multiple] option {
  padding: 0.375rem 0.75rem;
  margin: 0.125rem 0;
  border-radius: 0.25rem;
}

select[multiple] option:checked {
  background: linear-gradient(0deg, #007bff 0%, #0056b3 100%);
  color: white;
  font-weight: 500;
}

select[multiple]:focus {
  border-color: #80bdff;
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}

/* 表单标签样式 */
.form-label {
  font-weight: 600;
  color: #495057;
}

/* 下拉选择框样式 */
.form-select {
  transition: all 0.2s ease;
}

.form-select:hover {
  border-color: #007bff;
}

.form-select:focus {
  border-color: #80bdff;
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}

/* 表单文本提示样式 */
.form-text {
  font-size: 0.875rem;
  color: #6c757d;
  margin-top: 0.25rem;
}

/* 创建成功卡片动画 */
.card .card-header.bg-success {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%) !important;
}

/* 选择按钮样式 */
.btn-outline-primary {
  transition: all 0.3s ease;
}

.btn-outline-primary:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(0, 123, 255, 0.2);
}
</style>
