<template>
  <div class="ipfs-upload">
    <h2><i class="bi bi-cloud-upload me-2"></i>IPFS文件管理</h2>

    <div class="row mt-4">
      <!-- 上传区域 -->
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-upload me-2"></i>文件上传</h5>
          </div>
          <div class="card-body">
            <!-- 文件上传方式选择 -->
            <div class="mb-3">
              <div class="btn-group w-100" role="group">
                <input type="radio" class="btn-check" name="uploadType" id="fileUpload" v-model="uploadType" value="file" checked>
                <label class="btn btn-outline-primary" for="fileUpload">
                  <i class="bi bi-file-earmark me-1"></i>文件上传
                </label>

                <input type="radio" class="btn-check" name="uploadType" id="textUpload" v-model="uploadType" value="text">
                <label class="btn btn-outline-primary" for="textUpload">
                  <i class="bi bi-file-text me-1"></i>文本上传
                </label>
              </div>
            </div>

            <!-- 文件上传区域 -->
            <div v-if="uploadType === 'file'" class="mb-3">
              <label for="fileInput" class="form-label">选择文件</label>
              <input 
                type="file" 
                class="form-control" 
                id="fileInput" 
                ref="fileInput"
                @change="handleFileSelect"
                multiple
              >
              <div class="form-text">支持多文件上传，支持所有文件类型</div>
              
              <!-- 拖拽上传区域 -->
              <div 
                class="upload-drop-zone mt-3"
                :class="{ 'dragover': isDragOver }"
                @drop="handleDrop"
                @dragover="handleDragOver"
                @dragleave="handleDragLeave"
              >
                <div class="text-center py-4">
                  <i class="bi bi-cloud-upload display-4 text-muted"></i>
                  <p class="mt-2 mb-0">拖拽文件到此处或点击上方选择文件</p>
                  <small class="text-muted">支持批量上传</small>
                </div>
              </div>
            </div>

            <!-- 文本上传区域 -->
            <div v-if="uploadType === 'text'" class="mb-3">
              <label for="textContent" class="form-label">文本内容</label>
              <textarea 
                class="form-control" 
                id="textContent" 
                v-model="textContent" 
                rows="6"
                placeholder="输入要上传到IPFS的文本内容..."
              ></textarea>
              <div class="form-text">直接输入文本内容上传到IPFS</div>
              
              <div class="mt-2">
                <label for="fileName" class="form-label">文件名（可选）</label>
                <input 
                  type="text" 
                  class="form-control" 
                  id="fileName" 
                  v-model="fileName"
                  placeholder="例如：document.txt"
                >
              </div>
            </div>

            <!-- 上传按钮 -->
            <div class="d-flex gap-2">
              <button 
                type="button" 
                class="btn btn-primary" 
                @click="uploadToIPFS"
                :disabled="isUploading || !canUpload"
              >
                <span v-if="isUploading" class="spinner-border spinner-border-sm me-1"></span>
                <i v-else class="bi bi-cloud-upload me-1"></i>
                {{ isUploading ? '上传中...' : '上传到IPFS' }}
              </button>
              
              <button 
                type="button" 
                class="btn btn-outline-secondary" 
                @click="clearUpload"
                :disabled="isUploading"
              >
                <i class="bi bi-arrow-clockwise me-1"></i>清空
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 结果展示区域 -->
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-list-ul me-2"></i>上传结果</h5>
          </div>
          <div class="card-body">
            <div v-if="uploadResults.length === 0" class="text-center py-4 text-muted">
              <i class="bi bi-cloud" style="font-size: 3rem;"></i>
              <p class="mt-2">选择文件或输入文本后点击上传</p>
            </div>

            <div v-else>
              <div 
                v-for="(result, index) in uploadResults" 
                :key="index"
                class="result-item mb-3 p-3 border rounded"
              >
                <div class="d-flex justify-content-between align-items-start">
                  <div class="flex-grow-1">
                    <h6 class="mb-1">
                      <i class="bi bi-file-earmark me-1"></i>
                      {{ result.filename }}
                    </h6>
                    <div class="mb-2">
                      <small class="text-muted">
                        <i class="bi bi-calendar me-1"></i>{{ formatDate(result.uploadTime) }}
                        <span class="ms-2">
                          <i class="bi bi-file-binary me-1"></i>{{ formatFileSize(result.size) }}
                        </span>
                      </small>
                    </div>
                    
                    <div class="mb-2">
                      <label class="form-label small"><strong>IPFS Hash:</strong></label>
                      <div class="input-group input-group-sm">
                        <input 
                          type="text" 
                          class="form-control font-monospace" 
                          :value="result.hash" 
                          readonly
                          @click="$event.target.select()"
                        >
                        <button 
                          class="btn btn-outline-secondary" 
                          @click="copyToClipboard(result.hash)"
                        >
                          <i class="bi bi-clipboard"></i>
                        </button>
                      </div>
                    </div>

                    <div class="mb-2">
                      <label class="form-label small"><strong>IPFS URL:</strong></label>
                      <div class="input-group input-group-sm">
                        <input 
                          type="text" 
                          class="form-control" 
                          :value="result.url" 
                          readonly
                          @click="$event.target.select()"
                        >
                        <button 
                          class="btn btn-outline-secondary" 
                          @click="copyToClipboard(result.url)"
                        >
                          <i class="bi bi-clipboard"></i>
                        </button>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="d-flex gap-2 mt-2">
                  <button 
                    class="btn btn-sm btn-outline-primary" 
                    @click="viewOnIPFS(result.hash)"
                  >
                    <i class="bi bi-box-arrow-up-right me-1"></i>在IPFS查看
                  </button>
                  <button 
                    class="btn btn-sm btn-outline-info" 
                    @click="downloadFile(result)"
                  >
                    <i class="bi bi-download me-1"></i>下载文件
                  </button>
                  <button 
                    class="btn btn-sm btn-outline-danger" 
                    @click="removeResult(index)"
                  >
                    <i class="bi bi-trash me-1"></i>移除
                  </button>
                </div>
              </div>

              <div class="text-center mt-3">
                <button 
                  class="btn btn-outline-warning btn-sm" 
                  @click="clearAllResults"
                >
                  <i class="bi bi-trash me-1"></i>清空所有结果
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Hash下载功能 -->
    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0"><i class="bi bi-download me-2"></i>通过Hash获取文件</h5>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-md-8">
            <div class="mb-3">
              <label for="hashInput" class="form-label">IPFS Hash</label>
              <div class="input-group">
                <input 
                  type="text" 
                  class="form-control font-monospace" 
                  id="hashInput"
                  v-model="hashInput"
                  placeholder="输入IPFS Hash值，例如：QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG"
                  :disabled="isDownloading"
                >
                <button 
                  class="btn btn-outline-primary" 
                  @click="getFromIPFS"
                  :disabled="!hashInput.trim() || isDownloading"
                >
                  <span v-if="isDownloading" class="spinner-border spinner-border-sm me-1"></span>
                  <i v-else class="bi bi-search me-1"></i>
                  {{ isDownloading ? '获取中...' : '获取内容' }}
                </button>
              </div>
              <div class="form-text">输入完整的IPFS Hash值来获取存储的内容</div>
            </div>

            <!-- 获取结果展示 -->
            <div v-if="hashResult" class="mt-4">
              <div class="alert alert-success">
                <h6><i class="bi bi-check-circle me-2"></i>获取成功！</h6>
                <div class="row">
                  <div class="col-md-6">
                    <p class="mb-1"><strong>Hash:</strong></p>
                    <div class="bg-light p-2 rounded text-monospace small">{{ hashResult.hash }}</div>
                  </div>
                  <div class="col-md-6">
                    <p class="mb-1"><strong>文件大小:</strong></p>
                    <div class="bg-light p-2 rounded small">{{ formatFileSize(hashResult.size) }}</div>
                  </div>
                </div>
                <div class="mt-2">
                  <p class="mb-1"><strong>获取时间:</strong></p>
                  <div class="bg-light p-2 rounded small">{{ formatDate(hashResult.fetchTime) }}</div>
                </div>
              </div>

              <!-- 内容预览 -->
              <div class="mb-3">
                <label class="form-label"><strong>文件内容预览:</strong></label>
                <div class="content-preview-container">
                  <!-- 文本内容 -->
                  <div v-if="hashResult.isText" class="content-preview">
                    <textarea 
                      class="form-control" 
                      rows="8" 
                      readonly 
                      v-model="hashResult.content"
                      @click="$event.target.select()"
                    ></textarea>
                  </div>
                  
                  <!-- 图片内容 -->
                  <div v-else-if="hashResult.isImage" class="content-preview text-center">
                    <img 
                      :src="hashResult.url" 
                      alt="IPFS图片" 
                      class="img-fluid rounded"
                      style="max-height: 400px; max-width: 100%;"
                      @error="handleImageError"
                    >
                  </div>
                  
                  <!-- 其他文件类型 -->
                  <div v-else class="content-preview">
                    <div class="alert alert-info">
                      <i class="bi bi-info-circle me-2"></i>
                      此文件类型不支持预览，但可以下载查看。
                      <br><small>文件类型: {{ hashResult.contentType || '未知' }}</small>
                    </div>
                  </div>
                </div>
              </div>

                             <!-- 操作按钮 -->
               <div class="d-flex gap-2 flex-wrap">
                 <button 
                   class="btn btn-primary btn-sm" 
                   @click="downloadHashFile"
                 >
                   <i class="bi bi-download me-1"></i>下载文件
                 </button>
                 <button 
                   class="btn btn-outline-primary btn-sm" 
                   @click="downloadViaAPI"
                   v-if="hashResult.downloadUrl"
                 >
                   <i class="bi bi-cloud-download me-1"></i>通过API下载
                 </button>
                 <button 
                   class="btn btn-outline-info btn-sm" 
                   @click="copyToClipboard(hashResult.content)"
                   v-if="hashResult.isText"
                 >
                   <i class="bi bi-clipboard me-1"></i>复制内容
                 </button>
                 <button 
                   class="btn btn-outline-secondary btn-sm" 
                   @click="viewOnIPFS(hashResult.hash)"
                 >
                   <i class="bi bi-box-arrow-up-right me-1"></i>在IPFS查看
                 </button>
                 <button 
                   class="btn btn-outline-warning btn-sm" 
                   @click="clearHashResult"
                 >
                   <i class="bi bi-x me-1"></i>清除结果
                 </button>
               </div>
            </div>
          </div>

          <div class="col-md-4">
            <div class="card border-info">
              <div class="card-header bg-info text-white">
                <h6 class="mb-0"><i class="bi bi-lightbulb me-1"></i>使用提示</h6>
              </div>
              <div class="card-body">
                <ul class="list-unstyled mb-0">
                  <li class="mb-2">
                    <i class="bi bi-check-circle text-success me-1"></i>
                    <small>支持获取任何IPFS存储的文件</small>
                  </li>
                  <li class="mb-2">
                    <i class="bi bi-check-circle text-success me-1"></i>
                    <small>自动识别文件类型并预览</small>
                  </li>
                  <li class="mb-2">
                    <i class="bi bi-check-circle text-success me-1"></i>
                    <small>支持文本、图片等格式预览</small>
                  </li>
                  <li class="mb-2">
                    <i class="bi bi-check-circle text-success me-1"></i>
                    <small>一键下载到本地</small>
                  </li>
                  <li>
                    <i class="bi bi-info-circle text-info me-1"></i>
                    <small>Hash值通常以Qm开头</small>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- IPFS网关信息 -->
    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0"><i class="bi bi-info-circle me-2"></i>IPFS访问说明</h5>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-md-6">
            <h6>支持的IPFS网关：</h6>
            <ul class="list-unstyled">
              <li><i class="bi bi-link me-1"></i>https://dweb.link/ipfs/</li>
              <li><i class="bi bi-link me-1"></i>https://cloudflare-ipfs.com/ipfs/</li>
              <li><i class="bi bi-link me-1"></i>https://gateway.pinata.cloud/ipfs/</li>
              <li><i class="bi bi-link me-1"></i>https://ipfs.io/ipfs/</li>
            </ul>
          </div>
          <div class="col-md-6">
            <h6>使用说明：</h6>
            <ul class="list-unstyled">
              <li><i class="bi bi-check me-1"></i>支持任意文件类型上传</li>
              <li><i class="bi bi-check me-1"></i>支持文本直接输入上传</li>
              <li><i class="bi bi-check me-1"></i>支持拖拽批量上传</li>
              <li><i class="bi bi-check me-1"></i>自动生成IPFS Hash和访问链接</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, computed } from 'vue'
import { useStore } from 'vuex'
import ipfsService from '@/services/ipfsService'

export default {
  name: 'IPFSUpload',
  setup() {
    const store = useStore()

    const uploadType = ref('file')
    const textContent = ref('')
    const fileName = ref('')
    const isUploading = ref(false)
    const isDragOver = ref(false)
    const uploadResults = ref([])
    const fileInput = ref(null)

    // 添加文件选择状态
    const hasSelectedFiles = ref(false)

    // Hash下载相关状态
    const hashInput = ref('')
    const isDownloading = ref(false)
    const hashResult = ref(null)

    // 计算是否可以上传
    const canUpload = computed(() => {
      if (uploadType.value === 'file') {
        return hasSelectedFiles.value
      } else {
        return textContent.value.trim().length > 0
      }
    })

    // 处理文件选择
    const handleFileSelect = (event) => {
      const files = Array.from(event.target.files)
      hasSelectedFiles.value = files.length > 0
      console.log('选择了文件:', files.map(f => f.name))
    }

    // 处理拖拽
    const handleDrop = (event) => {
      event.preventDefault()
      event.stopPropagation()
      isDragOver.value = false

      const files = Array.from(event.dataTransfer.files)
      if (files.length > 0) {
        // 将文件设置到file input中
        const dt = new DataTransfer()
        files.forEach(file => dt.items.add(file))
        fileInput.value.files = dt.files
        hasSelectedFiles.value = true
        
        console.log('拖拽上传文件:', files.map(f => f.name))
      }
    }

    const handleDragOver = (event) => {
      event.preventDefault()
      event.stopPropagation()
      isDragOver.value = true
    }

    const handleDragLeave = (event) => {
      event.preventDefault()
      event.stopPropagation()
      isDragOver.value = false
    }

    // 读取文件内容 - 修复版本：正确处理二进制文件
    const readFileAsText = (file) => {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.onload = () => resolve(reader.result)
        reader.onerror = reject
        reader.readAsText(file)
      })
    }

    // 读取文件为ArrayBuffer - 新增函数
    const readFileAsArrayBuffer = (file) => {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.onload = () => resolve(reader.result)
        reader.onerror = reject
        reader.readAsArrayBuffer(file)
      })
    }

    // 检测文件是否为文本文件
    const isTextFile = (file) => {
      const textTypes = [
        'text/',
        'application/json',
        'application/xml',
        'application/javascript',
        'application/typescript'
      ]
      
      return textTypes.some(type => file.type.startsWith(type)) ||
             ['.txt', '.json', '.xml', '.html', '.css', '.js', '.ts', '.md', '.csv'].some(ext => 
               file.name.toLowerCase().endsWith(ext)
             )
    }

    // 将ArrayBuffer转换为Base64
    const arrayBufferToBase64 = (buffer) => {
      const bytes = new Uint8Array(buffer)
      let binary = ''
      for (let i = 0; i < bytes.byteLength; i++) {
        binary += String.fromCharCode(bytes[i])
      }
      return btoa(binary)
    }

    // 上传到IPFS - 修复版本
    const uploadToIPFS = async () => {
      console.log('uploadToIPFS 被调用')
      console.log('canUpload:', canUpload.value)
      console.log('uploadType:', uploadType.value)
      console.log('hasSelectedFiles:', hasSelectedFiles.value)
      console.log('textContent:', textContent.value)

      if (!canUpload.value) {
        console.log('canUpload 检查失败')
        store.dispatch('app/showError', '请选择文件或输入文本内容')
        return
      }

      isUploading.value = true
      console.log('开始上传，isUploading 设置为 true')
      
      try {
        if (uploadType.value === 'file') {
          // 文件上传 - 修复版本
          console.log('文件上传模式')
          if (!fileInput.value || !fileInput.value.files) {
            throw new Error('文件输入元素不可用')
          }
          
          const files = Array.from(fileInput.value.files)
          console.log('选择的文件:', files)
          
          for (const file of files) {
            console.log('处理文件:', file.name, '类型:', file.type, '大小:', file.size)
            
            let content
            let isBinary = false
            
            if (isTextFile(file)) {
              // 文本文件：按文本方式读取
              console.log('作为文本文件处理')
              content = await readFileAsText(file)
            } else {
              // 二进制文件：按ArrayBuffer读取并转换为Base64
              console.log('作为二进制文件处理')
              const arrayBuffer = await readFileAsArrayBuffer(file)
              content = arrayBufferToBase64(arrayBuffer)
              isBinary = true
            }
            
            console.log('文件内容读取完成，长度:', content.length, '是否为二进制:', isBinary)
            
            // 调用修改后的uploadToIPFS方法
            const result = await ipfsService.uploadToIPFS(content, file.name, isBinary)
            console.log('IPFS上传结果:', result)
            
            uploadResults.value.unshift({
              filename: file.name,
              hash: result.hash,
              url: result.url || `ipfs://${result.hash}`,
              size: file.size,
              uploadTime: new Date().toISOString(),
              content: content,
              displayContent: isBinary ? `[二进制文件 - ${file.type}]` : content,
              isBinary: isBinary,
              contentType: file.type
            })
          }
          
          store.dispatch('app/showSuccess', `成功上传 ${files.length} 个文件到IPFS`)
        } else {
          // 文本上传
          console.log('文本上传模式')
          const name = fileName.value.trim() || 'document.txt'
          console.log('文件名:', name)
          console.log('文本内容:', textContent.value)
          
          const result = await ipfsService.uploadToIPFS(textContent.value, name, false)
          console.log('IPFS上传结果:', result)
          
          uploadResults.value.unshift({
            filename: name,
            hash: result.hash,
            url: result.url || `ipfs://${result.hash}`,
            size: new Blob([textContent.value]).size,
            uploadTime: new Date().toISOString(),
            content: textContent.value,
            displayContent: textContent.value,
            isBinary: false,
            contentType: 'text/plain'
          })
          
          store.dispatch('app/showSuccess', '文本内容已成功上传到IPFS')
        }
      } catch (error) {
        console.error('上传到IPFS失败:', error)
        store.dispatch('app/showError', '上传失败: ' + error.message)
      } finally {
        isUploading.value = false
        console.log('上传完成，isUploading 设置为 false')
      }
    }

    // 清空上传
    const clearUpload = () => {
      if (uploadType.value === 'file') {
        if (fileInput.value) {
          fileInput.value.value = ''
        }
        hasSelectedFiles.value = false
      } else {
        textContent.value = ''
        fileName.value = ''
      }
    }

    // 复制到剪贴板
    const copyToClipboard = async (text) => {
      try {
        await navigator.clipboard.writeText(text)
        store.dispatch('app/showSuccess', '已复制到剪贴板')
      } catch (error) {
        console.error('复制失败:', error)
        store.dispatch('app/showError', '复制失败，请手动复制')
      }
    }

    // 在IPFS查看
    const viewOnIPFS = (hash) => {
      const url = ipfsService.getIPFSUrl(hash)
      window.open(url, '_blank')
    }

    // 下载文件 - 修复版本：正确处理二进制和文本数据
    const downloadFile = (result) => {
      try {
        console.log('下载文件:', result.filename, '是否为二进制:', result.isBinary, '内容类型:', result.contentType)
        
        let blob
        
        if (result.isBinary) {
          // 二进制文件处理
          if (!result.content || typeof result.content !== 'string') {
            store.dispatch('app/showError', '无法下载二进制文件：内容数据无效')
            return
          }
          
          try {
            // Base64解码
            console.log('解码Base64数据，长度:', result.content.length)
            const binaryString = atob(result.content)
            const bytes = new Uint8Array(binaryString.length)
            for (let i = 0; i < binaryString.length; i++) {
              bytes[i] = binaryString.charCodeAt(i)
            }
            
            blob = new Blob([bytes], { 
              type: result.contentType || 'application/octet-stream' 
            })
            
            console.log('二进制Blob创建成功，大小:', blob.size)
          } catch (error) {
            console.error('Base64解码失败:', error)
            store.dispatch('app/showError', '无法下载文件：二进制数据解码失败')
            return
          }
        } else {
          // 文本文件：直接使用内容
          blob = new Blob([result.content], { 
            type: result.contentType || 'text/plain' 
          })
          console.log('文本Blob创建成功，大小:', blob.size)
        }

        // 验证Blob大小
        if (blob.size === 0) {
          store.dispatch('app/showError', '无法下载文件：文件内容为空')
          return
        }

        const url = URL.createObjectURL(blob)
        const a = document.createElement('a')
        a.href = url
        a.download = result.filename
        a.style.display = 'none'
        document.body.appendChild(a)
        a.click()
        document.body.removeChild(a)
        URL.revokeObjectURL(url)
        
        store.dispatch('app/showSuccess', `文件下载成功: ${result.filename} (${formatFileSize(blob.size)})`)
        console.log('文件下载成功:', result.filename, '实际大小:', blob.size)
      } catch (error) {
        console.error('下载失败:', error)
        store.dispatch('app/showError', '下载失败: ' + error.message)
      }
    }

    // 移除结果
    const removeResult = (index) => {
      uploadResults.value.splice(index, 1)
    }

    // 清空所有结果
    const clearAllResults = () => {
      if (confirm('确定要清空所有上传结果吗？')) {
        uploadResults.value = []
        store.dispatch('app/showSuccess', '已清空所有结果')
      }
    }

    // 格式化日期
    const formatDate = (dateString) => {
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    // 格式化文件大小
    const formatFileSize = (bytes) => {
      if (bytes === 0) return '0 Bytes'
      const k = 1024
      const sizes = ['Bytes', 'KB', 'MB', 'GB']
      const i = Math.floor(Math.log(bytes) / Math.log(k))
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
    }

    // 通过Hash获取IPFS内容
    const getFromIPFS = async () => {
      const hash = hashInput.value.trim()
      if (!hash) {
        store.dispatch('app/showError', '请输入IPFS Hash值')
        return
      }

      // 验证Hash格式（基本验证）
      if (hash.length < 40 || (!hash.startsWith('Qm') && !hash.startsWith('bafy'))) {
        store.dispatch('app/showError', 'Hash格式不正确，请输入有效的IPFS Hash')
        return
      }

      isDownloading.value = true
      try {
        console.log('开始获取IPFS内容，Hash:', hash)
        
        // 使用优化的ipfsService获取内容
        const result = await ipfsService.getFromIPFSOptimized(hash)
        
        hashResult.value = {
          hash: result.hash,
          content: result.content,
          url: result.url,
          size: result.size,
          isText: result.isText,
          isImage: result.isImage,
          contentType: result.contentType,
          fetchTime: new Date().toISOString(),
          filename: `ipfs_${hash.substring(0, 8)}.txt`,
          downloadUrl: result.downloadUrl,
          source: result.source
        }

        store.dispatch('app/showSuccess', `成功获取IPFS内容! Hash: ${hash.substring(0, 12)}... (来源: ${result.source})`)
        console.log('IPFS内容获取成功:', hashResult.value)

      } catch (error) {
        console.error('获取IPFS内容失败:', error)
        store.dispatch('app/showError', '获取失败: ' + error.message)
      } finally {
        isDownloading.value = false
      }
    }

    // 下载通过Hash获取的文件 - 修复版本
    const downloadHashFile = () => {
      if (!hashResult.value) return

      try {
        const result = hashResult.value
        
        // 根据内容类型设置合适的文件扩展名
        let filename = `ipfs_${result.hash.substring(0, 8)}`
        
        if (result.isText) {
          if (result.contentType === 'application/json') {
            filename += '.json'
          } else if (result.contentType === 'text/html') {
            filename += '.html'
          } else {
            filename += '.txt'
          }
        } else if (result.isImage) {
          if (result.contentType === 'image/jpeg') filename += '.jpg'
          else if (result.contentType === 'image/png') filename += '.png'
          else if (result.contentType === 'image/gif') filename += '.gif'
          else if (result.contentType === 'image/svg+xml') filename += '.svg'
          else if (result.contentType === 'image/webp') filename += '.webp'
          else filename += '.img'
        } else {
          filename += '.bin'
        }

        // 正确处理二进制和文本数据
        let blob
        if (result.isBinary && result.content instanceof ArrayBuffer) {
          // 二进制数据：直接使用ArrayBuffer
          blob = new Blob([result.content], { type: result.contentType })
        } else if (result.isBinary && typeof result.content === 'string') {
          // 如果是损坏的字符串表示的二进制数据，尝试修复
          console.warn('检测到可能损坏的二进制数据，尝试修复...')
          try {
            // 尝试将字符串转换回二进制
            const binaryString = result.content
            const bytes = new Uint8Array(binaryString.length)
            for (let i = 0; i < binaryString.length; i++) {
              bytes[i] = binaryString.charCodeAt(i)
            }
            blob = new Blob([bytes], { type: result.contentType })
          } catch (error) {
            console.error('二进制数据修复失败:', error)
            // 如果修复失败，仍然尝试下载
            blob = new Blob([result.content], { type: result.contentType })
          }
        } else {
          // 文本数据
          blob = new Blob([result.content], { type: result.contentType })
        }

        const url = URL.createObjectURL(blob)
        const a = document.createElement('a')
        a.href = url
        a.download = filename
        document.body.appendChild(a)
        a.click()
        document.body.removeChild(a)
        URL.revokeObjectURL(url)
        
        store.dispatch('app/showSuccess', `文件 ${filename} 下载已开始`)
      } catch (error) {
        console.error('下载失败:', error)
        store.dispatch('app/showError', '下载失败: ' + error.message)
      }
    }

    // 通过API下载文件
    const downloadViaAPI = () => {
      if (!hashResult.value || !hashResult.value.downloadUrl) return

      try {
        const downloadUrl = hashResult.value.downloadUrl
        
        // 创建隐藏的链接并触发下载
        const a = document.createElement('a')
        a.href = downloadUrl
        a.download = '' // 让浏览器使用服务器提供的文件名
        a.style.display = 'none'
        document.body.appendChild(a)
        a.click()
        document.body.removeChild(a)
        
        store.dispatch('app/showSuccess', '通过API下载已开始')
      } catch (error) {
        console.error('API下载失败:', error)
        store.dispatch('app/showError', 'API下载失败: ' + error.message)
      }
    }

    // 清除Hash获取结果
    const clearHashResult = () => {
      hashResult.value = null
      hashInput.value = ''
      store.dispatch('app/showSuccess', '已清除获取结果')
    }

    // 处理图片加载错误
    const handleImageError = (event) => {
      console.error('图片加载失败:', event.target.src)
      if (!event.target.src.includes('placeholder')) {
        // 使用占位图片
        event.target.src = 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22120%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2214%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3E%E5%9B%BE%E7%89%87%E5%8A%A0%E8%BD%BD%E5%A4%B1%E8%B4%A5%3C%2Ftext%3E%3C%2Fsvg%3E'
        event.target.onerror = null
      }
    }

    return {
      uploadType,
      textContent,
      fileName,
      isUploading,
      isDragOver,
      uploadResults,
      fileInput,
      hasSelectedFiles,
      canUpload,
      hashInput,
      isDownloading,
      hashResult,
      handleFileSelect,
      handleDrop,
      handleDragOver,
      handleDragLeave,
      uploadToIPFS,
      clearUpload,
      copyToClipboard,
      viewOnIPFS,
      downloadFile,
      removeResult,
      clearAllResults,
      formatDate,
      formatFileSize,
      getFromIPFS,
      downloadHashFile,
      downloadViaAPI,
      clearHashResult,
      handleImageError
    }
  }
}
</script>

<style scoped>
.upload-drop-zone {
  border: 2px dashed #dee2e6;
  border-radius: 0.5rem;
  transition: all 0.3s ease;
  cursor: pointer;
}

.upload-drop-zone:hover,
.upload-drop-zone.dragover {
  border-color: #0d6efd;
  background-color: rgba(13, 110, 253, 0.05);
}

.result-item {
  background-color: #f8f9fa;
  transition: all 0.2s ease;
}

.result-item:hover {
  background-color: #e9ecef;
}

.font-monospace {
  font-family: 'Courier New', monospace;
  font-size: 0.9rem;
}

.btn-group .btn-check:checked + .btn {
  background-color: var(--bs-primary);
  border-color: var(--bs-primary);
  color: white;
}

.content-preview-container {
  border: 1px solid #dee2e6;
  border-radius: 0.5rem;
  overflow: hidden;
}

.content-preview {
  background-color: #f8f9fa;
  border-radius: 0.5rem;
}

.content-preview textarea {
  border: none;
  background-color: transparent;
  resize: vertical;
}

.content-preview img {
  border: 1px solid #dee2e6;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.card.border-info {
  border-width: 2px;
}

.bg-info {
  background-color: #0dcaf0 !important;
}

.text-monospace {
  font-family: 'Courier New', monospace;
  font-size: 0.9rem;
}
</style> 