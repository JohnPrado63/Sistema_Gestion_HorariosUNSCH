<template>
  <div class="user-menu-container" v-if="user">
    <button class="user-button" @click="toggleMenu" ref="buttonRef">
      <div class="user-avatar">
        {{ userInitials }}
      </div>
      <div class="user-info">
        <span class="user-name">{{ user.nombre }}</span>
        <span class="user-role">{{ formatRole(user.rol) }}</span>
      </div>
      <svg class="chevron" :class="{ 'open': isOpen }" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <polyline points="6 9 12 15 18 9"/>
      </svg>
    </button>

    <Transition name="dropdown">
      <div v-if="isOpen" class="dropdown-menu">
        <div class="dropdown-header">
          <div class="user-avatar-lg">
            {{ userInitials }}
          </div>
          <div class="user-details">
            <span class="user-name-lg">{{ user.nombre }}</span>
            <span class="user-email">{{ user.email }}</span>
          </div>
        </div>

        <div class="dropdown-divider"></div>

        <button class="dropdown-item" @click="goToProfile">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
            <circle cx="12" cy="7" r="4"/>
          </svg>
          Mi Perfil
        </button>

        <div class="dropdown-divider"></div>

        <button class="dropdown-item logout" @click="handleLogout">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
            <polyline points="16 17 21 12 16 7"/>
            <line x1="21" y1="12" x2="9" y2="12"/>
          </svg>
          Cerrar Sesión
        </button>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../context/AuthContext'

const router = useRouter()
const { user, logout } = useAuth()

const isOpen = ref(false)
const buttonRef = ref(null)

const userInitials = computed(() => {
  if (!user.value?.nombre) return '?'
  const parts = user.value.nombre.split(' ')
  if (parts.length >= 2) {
    return (parts[0][0] + parts[1][0]).toUpperCase()
  }
  return parts[0].substring(0, 2).toUpperCase()
})

const roleColors = {
  'ADMIN_TI': { bg: 'linear-gradient(135deg, #667eea, #764ba2)', text: 'Administrador TI' },
  'DGA': { bg: 'linear-gradient(135deg, #f093fb, #f5576c)', text: 'Director General Académico' },
  'DIRECTOR_ESCUELA': { bg: 'linear-gradient(135deg, #11998e, #38ef7d)', text: 'Director de Escuela' },
  'JEFE_DEPTO': { bg: 'linear-gradient(135deg, #fc466b, #3f5efb)', text: 'Jefe de Departamento' },
  'COORDINADOR': { bg: 'linear-gradient(135deg, #f6ad55, #dd6b20)', text: 'Coordinador' },
  'DOCENTE': { bg: 'linear-gradient(135deg, #667eea, #764ba2)', text: 'Docente' },
  'ESTUDIANTE': { bg: 'linear-gradient(135deg, #3b82f6, #8b5cf6)', text: 'Estudiante' }
}

function formatRole(rol) {
  return roleColors[rol]?.text || rol
}

function toggleMenu() {
  isOpen.value = !isOpen.value
}

function handleClickOutside(event) {
  if (buttonRef.value && !buttonRef.value.contains(event.target)) {
    isOpen.value = false
  }
}

function goToProfile() {
  isOpen.value = false
  console.log('Go to profile')
}

async function handleLogout() {
  isOpen.value = false
  await logout()
  router.push('/login')
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.user-menu-container {
  position: relative;
}

.user-button {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px 12px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s;
  color: white;
}

.user-button:hover {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
}

.user-avatar {
  width: 36px;
  height: 36px;
  background: linear-gradient(135deg, #06b6d4, #0ea5e9);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.85rem;
  font-weight: 700;
  color: white;
}

.user-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-align: left;
}

.user-name {
  font-size: 0.85rem;
  font-weight: 600;
  color: white;
  max-width: 120px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-role {
  font-size: 0.7rem;
  color: rgba(255, 255, 255, 0.7);
}

.chevron {
  transition: transform 0.3s;
  color: rgba(255, 255, 255, 0.7);
}

.chevron.open {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  min-width: 280px;
  background: white;
  border-radius: 16px;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(0, 0, 0, 0.05);
  overflow: hidden;
  z-index: 1000;
}

.dropdown-header {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 20px;
  background: linear-gradient(135deg, #667eea15, #764ba215);
}

.user-avatar-lg {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.1rem;
  font-weight: 700;
  color: white;
}

.user-details {
  display: flex;
  flex-direction: column;
}

.user-name-lg {
  font-size: 1rem;
  font-weight: 700;
  color: #1e293b;
}

.user-email {
  font-size: 0.8rem;
  color: #64748b;
}

.dropdown-divider {
  height: 1px;
  background: #e2e8f0;
  margin: 0;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
  padding: 14px 20px;
  background: none;
  border: none;
  font-size: 0.9rem;
  color: #334155;
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
}

.dropdown-item:hover {
  background: #f8fafc;
  color: #1e293b;
}

.dropdown-item.logout {
  color: #dc2626;
}

.dropdown-item.logout:hover {
  background: #fef2f2;
}

.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.2s ease;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>
