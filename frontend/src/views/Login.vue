<template>
  <div class="login-page">
    <div class="bg-effects">
      <div class="orb orb-1"></div>
      <div class="orb orb-2"></div>
      <div class="orb orb-3"></div>
      <div class="grid-lines"></div>
    </div>

    <div class="login-container">
      <div class="login-card">
        <div class="card-header">
          <div class="logo">
            <div class="logo-icon">
              <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
              </svg>
            </div>
            <div class="logo-text">
              <span class="logo-title">UNSCH</span>
              <span class="logo-subtitle">Horarios Académicos</span>
            </div>
          </div>
        </div>

        <div class="card-body">
          <div class="welcome-text">
            <h1>¡Bienvenido!</h1>
            <p>Ingresa tus credenciales para acceder al sistema</p>
          </div>

          <form @submit.prevent="handleLogin" class="login-form">
            <div class="form-group" :class="{ 'focused': emailFocused, 'error': errors.email }">
              <div class="input-wrapper">
                <div class="input-icon">
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
                    <polyline points="22,6 12,13 2,6"/>
                  </svg>
                </div>
                <input
                  type="email"
                  v-model="email"
                  @focus="emailFocused = true"
                  @blur="emailFocused = false"
                  placeholder="Correo electrónico"
                  class="form-input"
                  required
                />
              </div>
              <span v-if="errors.email" class="error-text">{{ errors.email }}</span>
            </div>

            <div class="form-group" :class="{ 'focused': passwordFocused, 'error': errors.password }">
              <div class="input-wrapper">
                <div class="input-icon">
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                    <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                  </svg>
                </div>
                <input
                  :type="showPassword ? 'text' : 'password'"
                  v-model="password"
                  @focus="passwordFocused = true"
                  @blur="passwordFocused = false"
                  placeholder="Contraseña"
                  class="form-input"
                  required
                />
                <button type="button" class="toggle-password" @click="showPassword = !showPassword">
                  <svg v-if="!showPassword" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                    <circle cx="12" cy="12" r="3"/>
                  </svg>
                  <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"/>
                    <line x1="1" y1="1" x2="23" y2="23"/>
                  </svg>
                </button>
              </div>
              <span v-if="errors.password" class="error-text">{{ errors.password }}</span>
            </div>

            <div v-if="loginError" class="error-banner" :class="{ 'shake': shakeError }">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <line x1="12" y1="8" x2="12" y2="12"/>
                <line x1="12" y1="16" x2="12.01" y2="16"/>
              </svg>
              <span>{{ loginError }}</span>
            </div>

            <button type="submit" class="btn-login" :disabled="loading">
              <span v-if="loading" class="spinner"></span>
              <span v-else>
                Iniciar Sesión
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="5" y1="12" x2="19" y2="12"/>
                  <polyline points="12 5 19 12 12 19"/>
                </svg>
              </span>
            </button>
          </form>

          <div class="card-footer">
            <p>¿Olvidaste tu contraseña?</p>
            <p class="demo-hint">
              Demo: <strong>admin@unsch.edu.pe</strong> / <strong>admin123</strong>
            </p>
          </div>
        </div>
      </div>

      <div class="brand-footer">
        <p>Universidad Nacional de San Cristóbal de Huamanga</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../context/AuthContext'

const router = useRouter()
const { login } = useAuth()

const email = ref('')
const password = ref('')
const showPassword = ref(false)
const loading = ref(false)
const loginError = ref('')
const shakeError = ref(false)
const emailFocused = ref(false)
const passwordFocused = ref(false)
const errors = reactive({ email: '', password: '' })

async function handleLogin() {
  errors.email = ''
  errors.password = ''
  loginError.value = ''

  if (!email.value) {
    errors.email = 'El correo es requerido'
    return
  }
  if (!password.value) {
    errors.password = 'La contraseña es requerida'
    return
  }

  loading.value = true

  try {
    await login(email.value, password.value)
    router.push('/dashboard')
  } catch (error) {
    loginError.value = error.message
    shakeError.value = true
    setTimeout(() => {
      shakeError.value = false
    }, 600)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #0f172a 0%, #1e3a5f 50%, #0c1929 100%);
  position: relative;
  overflow: hidden;
}

.bg-effects {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.4;
  animation: float 20s ease-in-out infinite;
}

.orb-1 {
  width: 600px;
  height: 600px;
  background: radial-gradient(circle, #06b6d4 0%, transparent 70%);
  top: -200px;
  left: -200px;
  animation-delay: 0s;
}

.orb-2 {
  width: 500px;
  height: 500px;
  background: radial-gradient(circle, #0891b2 0%, transparent 70%);
  bottom: -150px;
  right: -150px;
  animation-delay: -7s;
}

.orb-3 {
  width: 400px;
  height: 400px;
  background: radial-gradient(circle, #0ea5e9 0%, transparent 70%);
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  animation-delay: -14s;
}

@keyframes float {
  0%, 100% { transform: translate(0, 0) scale(1); }
  25% { transform: translate(50px, -50px) scale(1.1); }
  50% { transform: translate(-30px, 30px) scale(0.95); }
  75% { transform: translate(-50px, -30px) scale(1.05); }
}

.grid-lines {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(6, 182, 212, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(6, 182, 212, 0.03) 1px, transparent 1px);
  background-size: 60px 60px;
}

.login-container {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 480px;
  padding: 20px;
}

.login-card {
  background: rgba(15, 23, 42, 0.8);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  border: 1px solid rgba(6, 182, 212, 0.2);
  box-shadow:
    0 25px 50px -12px rgba(0, 0, 0, 0.5),
    0 0 0 1px rgba(6, 182, 212, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.05);
  overflow: hidden;
  animation: slideUp 0.6s ease-out;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.card-header {
  padding: 32px 32px 0;
}

.logo {
  display: flex;
  align-items: center;
  gap: 16px;
}

.logo-icon {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, #06b6d4 0%, #0ea5e9 100%);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  box-shadow: 0 8px 24px rgba(6, 182, 212, 0.3);
}

.logo-text {
  display: flex;
  flex-direction: column;
}

.logo-title {
  font-size: 1.75rem;
  font-weight: 800;
  color: white;
  letter-spacing: 2px;
}

.logo-subtitle {
  font-size: 0.85rem;
  color: rgba(255, 255, 255, 0.6);
}

.card-body {
  padding: 32px;
}

.welcome-text {
  text-align: center;
  margin-bottom: 32px;
}

.welcome-text h1 {
  font-size: 1.75rem;
  font-weight: 700;
  color: white;
  margin: 0 0 8px;
}

.welcome-text p {
  color: rgba(255, 255, 255, 0.6);
  margin: 0;
  font-size: 0.95rem;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  position: relative;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.input-icon {
  position: absolute;
  left: 16px;
  color: rgba(255, 255, 255, 0.4);
  transition: color 0.3s;
  z-index: 1;
}

.form-group.focused .input-icon {
  color: #06b6d4;
}

.form-input {
  width: 100%;
  padding: 16px 50px 16px 50px;
  background: rgba(255, 255, 255, 0.05);
  border: 2px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  color: white;
  font-size: 1rem;
  transition: all 0.3s;
  box-sizing: border-box;
}

.form-input::placeholder {
  color: rgba(255, 255, 255, 0.4);
}

.form-input:focus {
  outline: none;
  border-color: #06b6d4;
  background: rgba(6, 182, 212, 0.05);
  box-shadow: 0 0 0 4px rgba(6, 182, 212, 0.1);
}

.form-group.error .form-input {
  border-color: #ef4444;
}

.toggle-password {
  position: absolute;
  right: 16px;
  background: none;
  border: none;
  color: rgba(255, 255, 255, 0.4);
  cursor: pointer;
  padding: 4px;
  transition: color 0.3s;
}

.toggle-password:hover {
  color: white;
}

.error-text {
  display: block;
  margin-top: 6px;
  font-size: 0.8rem;
  color: #ef4444;
}

.error-banner {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 16px;
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 12px;
  color: #ef4444;
  font-size: 0.9rem;
  animation: fadeIn 0.3s;
}

.error-banner.shake {
  animation: shake 0.5s ease-in-out;
}

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  20% { transform: translateX(-10px); }
  40% { transform: translateX(10px); }
  60% { transform: translateX(-10px); }
  80% { transform: translateX(10px); }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.btn-login {
  width: 100%;
  padding: 16px;
  background: linear-gradient(135deg, #06b6d4 0%, #0ea5e9 100%);
  border: none;
  border-radius: 12px;
  color: white;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  box-shadow: 0 8px 24px rgba(6, 182, 212, 0.3);
}

.btn-login:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 12px 32px rgba(6, 182, 212, 0.4);
}

.btn-login:active:not(:disabled) {
  transform: translateY(0);
}

.btn-login:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.card-footer {
  margin-top: 24px;
  text-align: center;
  padding-top: 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.card-footer p {
  color: rgba(255, 255, 255, 0.5);
  margin: 0;
  font-size: 0.85rem;
}

.demo-hint {
  margin-top: 12px !important;
  padding: 12px;
  background: rgba(6, 182, 212, 0.1);
  border-radius: 8px;
  font-size: 0.8rem !important;
  color: rgba(255, 255, 255, 0.7) !important;
}

.demo-hint strong {
  color: #06b6d4;
}

.brand-footer {
  text-align: center;
  margin-top: 24px;
}

.brand-footer p {
  color: rgba(255, 255, 255, 0.3);
  font-size: 0.8rem;
  margin: 0;
}

@media (max-width: 480px) {
  .login-container {
    padding: 10px;
  }

  .card-header,
  .card-body {
    padding: 24px;
  }

  .logo-icon {
    width: 50px;
    height: 50px;
  }

  .logo-title {
    font-size: 1.5rem;
  }
}
</style>
