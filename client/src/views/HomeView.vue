<script setup lang="ts">
import { ref } from 'vue'
</script>

<template>
  <div class="min-h-[calc(100vh-200px)] flex items-center justify-center">
    <div class="container mx-auto px-4 py-16">
      <!-- Hero Section -->
      <div class="text-center max-w-3xl mx-auto">
        <!-- Badge -->
        <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-emerald-500/10 border border-emerald-500/20 text-emerald-400 text-sm font-medium mb-8">
          <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
          Открыта регистрация
        </div>

        <!-- Title -->
        <h1 class="text-5xl md:text-6xl font-bold mb-6">
          <span class="bg-gradient-to-r from-emerald-400 via-cyan-400 to-emerald-400 bg-clip-text text-transparent">
            Играй в реальном времени
          </span>
        </h1>

        <!-- Subtitle -->
        <p class="text-xl text-gray-400 mb-10 leading-relaxed">
          Браузерная онлайн-игра на Go с WebSocket соединением.
          Минимальная задержка, максимальное удовольствие.
        </p>

        <!-- CTA Buttons -->
        <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
          <RouterLink
              to="/game"
              class="w-full sm:w-auto px-8 py-4 rounded-xl bg-gradient-to-r from-emerald-500 to-cyan-500 text-white font-semibold hover:from-emerald-600 hover:to-cyan-600 transition-all shadow-lg shadow-emerald-500/25 hover:shadow-emerald-500/40 hover:scale-105"
          >
            Начать игру
          </RouterLink>
          <RouterLink
              to="/about"
              class="w-full sm:w-auto px-8 py-4 rounded-xl border border-gray-700 text-gray-300 font-semibold hover:bg-gray-800 hover:border-gray-600 transition-all"
          >
            Узнать больше
          </RouterLink>
        </div>

        <!-- Features -->
        <div class="grid md:grid-cols-3 gap-6 mt-20">
          <div class="p-6 rounded-2xl bg-gray-800/50 border border-gray-700/50 hover:border-emerald-500/30 transition-colors">
            <div class="w-12 h-12 rounded-lg bg-emerald-500/10 flex items-center justify-center mb-4">
              <svg class="w-6 h-6 text-emerald-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
              </svg>
            </div>
            <h3 class="text-lg font-semibold mb-2">Мгновенно</h3>
            <p class="text-gray-400 text-sm">WebSocket соединение для минимальной задержки</p>
          </div>

          <div class="p-6 rounded-2xl bg-gray-800/50 border border-gray-700/50 hover:border-cyan-500/30 transition-colors">
            <div class="w-12 h-12 rounded-lg bg-cyan-500/10 flex items-center justify-center mb-4">
              <svg class="w-6 h-6 text-cyan-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
              </svg>
            </div>
            <h3 class="text-lg font-semibold mb-2">В браузере</h3>
            <p class="text-gray-400 text-sm">Никаких установок, играйте сразу</p>
          </div>

          <div class="p-6 rounded-2xl bg-gray-800/50 border border-gray-700/50 hover:border-purple-500/30 transition-colors">
            <div class="w-12 h-12 rounded-lg bg-purple-500/10 flex items-center justify-center mb-4">
              <svg class="w-6 h-6 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/>
              </svg>
            </div>
            <h3 class="text-lg font-semibold mb-2">Open Source</h3>
            <p class="text-gray-400 text-sm">Открытый код на GitHub</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>