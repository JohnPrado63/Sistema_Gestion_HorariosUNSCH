<template>
  <div class="app-layout" v-if="isAuthenticated">
    <aside class="sidebar">
      <div class="sidebar-header">
        <div class="sidebar-logo">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
          </svg>
          <div>
            UNSCH
            <span>Sistema de Horarios</span>
          </div>
        </div>
      </div>

      <nav class="sidebar-nav">
        <span class="nav-section">Principal</span>
        <router-link to="/dashboard" class="nav-item" active-class="active">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="3" width="7" height="7"/>
            <rect x="14" y="3" width="7" height="7"/>
            <rect x="14" y="14" width="7" height="7"/>
            <rect x="3" y="14" width="7" height="7"/>
          </svg>
          Dashboard
        </router-link>

        <span class="nav-section">Gestión</span>
        <router-link v-if="canAccess('catalogos')" to="/catalogos" class="nav-item" active-class="active">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
          </svg>
          Catálogos
        </router-link>
        <router-link v-if="canAccess('carga-academica')" to="/carga-academica" class="nav-item" active-class="active">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M22 12h-4l-3 9L9 3l-3 9H2"/>
          </svg>
          Carga Académica
        </router-link>
        <router-link v-if="canAccess('horarios')" to="/horarios" class="nav-item" active-class="active">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
            <line x1="16" y1="2" x2="16" y2="6"/>
            <line x1="8" y1="2" x2="8" y2="6"/>
            <line x1="3" y1="10" x2="21" y2="10"/>
          </svg>
          Horarios
        </router-link>
        <router-link v-if="canAccess('validaciones')" to="/validaciones" class="nav-item" active-class="active">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
            <polyline points="22 4 12 14.01 9 11.01"/>
          </svg>
          Validaciones
        </router-link>

        <span class="nav-section">Herramientas</span>
        <router-link v-if="canAccess('siige')" to="/siige" class="nav-item" active-class="active">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
            <polyline points="17 8 12 3 7 8"/>
            <line x1="12" y1="3" x2="12" y2="15"/>
          </svg>
          Importador SIIGE
        </router-link>
      </nav>
    </aside>

    <main class="main-content">
      <header class="topbar">
        <div class="topbar-left">
          <h1 class="page-title">{{ pageTitle }}</h1>
        </div>
        <div class="topbar-right">
          <UserMenu />
        </div>
      </header>
      <div class="content-area">
        <router-view />
      </div>
    </main>
  </div>

  <router-view v-else />
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import UserMenu from './components/UserMenu.vue'
import { useAuth } from './context/AuthContext'

const { user, isAuthenticated } = useAuth()
const route = useRoute()

const pageTitles = {
  '/dashboard': 'Panel de Control',
  '/catalogos': 'Catálogos Institucionales',
  '/carga-academica': 'Carga Académica',
  '/horarios': 'Gestión de Horarios',
  '/validaciones': 'Motor de Validaciones',
  '/siige': 'Importador SIIGE'
}

const pageTitle = computed(() => {
  return pageTitles[route.path] || 'Sistema de Horarios'
})

const roleRoutes = {
  'catalogos': ['ADMIN_TI', 'DGA'],
  'carga-academica': ['ADMIN_TI', 'DGA', 'DIRECTOR_ESCUELA', 'JEFE_DEPTO'],
  'horarios': ['ADMIN_TI', 'DGA', 'DIRECTOR_ESCUELA', 'JEFE_DEPTO', 'COORDINADOR'],
  'validaciones': ['ADMIN_TI', 'DGA'],
  'siige': ['ADMIN_TI', 'DGA']
}

const canAccess = (routeName) => {
  if (!user.value) return false
  const allowedRoles = roleRoutes[routeName]
  if (!allowedRoles) return true
  return allowedRoles.includes(user.value.rol)
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: #f8fafc;
  color: #1e293b;
}

.app-layout {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 260px;
  background: linear-gradient(180deg, #0f172a 0%, #1e293b 100%);
  color: white;
  display: flex;
  flex-direction: column;
  position: fixed;
  height: 100vh;
  overflow-y: auto;
  box-shadow: 4px 0 20px rgba(0, 0, 0, 0.1);
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.sidebar-logo {
  display: flex;
  align-items: center;
  gap: 12px;
}

.sidebar-logo svg {
  color: #06b6d4;
}

.sidebar-logo div {
  font-size: 1.1rem;
  font-weight: 700;
  display: flex;
  flex-direction: column;
}

.sidebar-logo span {
  font-size: 0.7rem;
  font-weight: 400;
  opacity: 0.6;
}

.sidebar-nav {
  padding: 16px 12px;
  flex: 1;
}

.nav-section {
  display: block;
  font-size: 0.7rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 1px;
  color: rgba(255, 255, 255, 0.4);
  padding: 12px 8px 8px;
  margin-top: 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  color: rgba(255, 255, 255, 0.7);
  text-decoration: none;
  border-radius: 10px;
  margin-bottom: 4px;
  transition: all 0.3s;
  font-size: 0.9rem;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.1);
  color: white;
}

.nav-item.active {
  background: linear-gradient(135deg, #06b6d4, #0ea5e9);
  color: white;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.nav-item svg {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
}

.main-content {
  flex: 1;
  margin-left: 260px;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.topbar {
  height: 70px;
  background: white;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  position: sticky;
  top: 0;
  z-index: 100;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

.topbar-left {
  display: flex;
  align-items: center;
}

.page-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: #1e293b;
  margin: 0;
}

.topbar-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.content-area {
  flex: 1;
  padding: 0;
}

::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.05);
}

::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.15);
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(0, 0, 0, 0.25);
}
</style>
