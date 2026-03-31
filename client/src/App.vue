<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="min-h-screen bg-gray-900 text-gray-100">
    <!-- Header -->
    <header class="border-b border-gray-800 bg-gray-900/80 backdrop-blur-sm sticky top-0 z-50">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <!-- Logo -->
          <RouterLink to="/" class="flex items-center gap-3 group">
            <div class="w-10 h-10 rounded-lg bg-gradient-to-br from-emerald-500 to-cyan-500 flex items-center justify-center group-hover:scale-105 transition-transform">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <span class="text-xl font-bold bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text text-transparent">
              WebSocket Game
            </span>
          </RouterLink>

          <!-- Navigation -->
          <nav class="flex items-center gap-6">
            <RouterLink
                to="/"
                class="text-gray-400 hover:text-emerald-400 transition-colors font-medium"
                active-class="text-emerald-400"
            >
              Главная
            </RouterLink>
            <RouterLink
                to="/game"
                class="text-gray-400 hover:text-emerald-400 transition-colors font-medium"
                active-class="text-emerald-400"
            >
              Игра
            </RouterLink>
            <RouterLink
                to="/about"
                class="text-gray-400 hover:text-emerald-400 transition-colors font-medium"
                active-class="text-emerald-400"
            >
              О проекте
            </RouterLink>
            <a
                href="https://github.com/emilversiontwo/WebSocketGame"
                target="_blank"
                class="text-gray-400 hover:text-white transition-colors"
                aria-label="GitHub"
            >
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.927 0-1.088.39-1.979 1.029-2.668-.102-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.69 1.028 1.58 1.028 2.668 0 3.825-2.339 4.673-4.566 4.92.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clip-rule="evenodd"/>
              </svg>
            </a>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main>
      <RouterView />
    </main>

    <!-- Footer -->
    <footer class="border-t border-gray-800 bg-gray-900 mt-auto">
      <div class="container mx-auto px-4 py-6">
        <div class="flex flex-col md:flex-row items-center justify-between gap-4">
          <p class="text-gray-500 text-sm">
            © 2026 WebSocket Game. Все права защищены.
          </p>
          <div class="flex items-center gap-4">
            <span class="flex items-center gap-2 text-sm text-gray-500">
              <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
              Сервер онлайн
            </span>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<style scoped>
/* Плавные переходы для RouterLink */
.router-link-active {
  position: relative;
}

.router-link-active::after {
  content: '';
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 100%;
  height: 2px;
  background: linear-gradient(to right, #10b981, #06b6d4);
  border-radius: 1px;
}
</style>