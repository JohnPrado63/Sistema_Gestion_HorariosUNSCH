const API_BASE = '/api/v1';

console.log('UNSCH frontend script loaded');

function el(q, parent=document) { return parent.querySelector(q); }
function els(q, parent=document) { return Array.from(parent.querySelectorAll(q)); }

function setStatus(msg){ const s = el('#appStatus'); if(s) s.textContent = msg; console.log('status:', msg); }

async function apiGET(path){
  const res = await fetch(API_BASE + path);
  if(!res.ok) throw new Error('API error '+res.status);
  return res.json();
}

function render(route){
  try{
    setStatus('navegando: '+route);
    const main = el('#main');
    main.innerHTML = '';
    if(route==='dashboard') return renderDashboard(main);
    if(route==='catalogs') return renderCatalogs(main);
    if(route==='validations') return renderValidations(main);
    if(route==='schedules') return renderSchedules(main);
    if(route==='siige') return renderSiige(main);
    renderDashboard(main);
  }catch(e){
    console.error('Render error', e);
    setStatus('Error en UI: '+e.message);
  }
}

async function renderDashboard(container){
  setStatus('mostrando dashboard');
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Resumen</h3><p class="small">Demo local del Motor de Horarios — APIs disponibles: <span class="code">/api/v1</span></p><div class="grid"><div class="card"><h4>Validaciones</h4><p class="small">RV-01..RV-09 disponibles en el backend.</p></div><div class="card"><h4>Importador SIIGE</h4><p class="small">Simulador en esta SPA para probar importes.</p></div></div>`;
  container.appendChild(card);
}

async function renderCatalogs(container){
  setStatus('mostrando catálogos');
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Catálogos</h3><p class="small">Cargar listados desde API para revisión.</p><div class="controls"><button type="button" class="btn" id="loadCatalogs">Cargar</button></div><div id="catalogResult" class="preview"></div>`;
  container.appendChild(card);
  el('#loadCatalogs').addEventListener('click',async ()=>{
    setStatus('cargando catálogos...');
    try{
      const facs = await apiGET('/facultades');
      const esc = await apiGET('/escuelas');
      document.getElementById('catalogResult').textContent = JSON.stringify({facultades:facs,escuelas:esc},null,2);
      setStatus('catálogos cargados');
    }catch(e){document.getElementById('catalogResult').textContent = 'Error: '+e.message; setStatus('error al cargar catálogos')}
  })
}

async function renderValidations(container){
  setStatus('mostrando validaciones');
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Validaciones</h3><p class="small">Enviar escenarios de prueba al backend.</p><div class="controls"><button type="button" class="btn" id="runSimple">Escenario simple</button><button type="button" class="btn secondary" id="runComplex">Escenario complejo</button></div><pre id="valResult" class="preview"></pre><div class="controls" style="margin-top:16px;"><button type="button" class="btn secondary" id="refreshScenarios">Actualizar escenarios</button></div><pre id="scenarioResult" class="preview">Cargando escenarios del motor...</pre>`;
  container.appendChild(card);
  el('#refreshScenarios').addEventListener('click', loadValidationScenarios);
  loadValidationScenarios();
  el('#runSimple').addEventListener('click', async ()=>{
    setStatus('ejecutando validación simple');
    try{
      const payload = {
        proposed: {
          id: 101,
          teacher_id: 10,
          school_id: 1,
          group_id: 1,
          room_id: 5,
          series_id: 1,
          day: 1,
          start_slot: 3,
          end_slot: 4,
          enrollment: 42,
          room_capacity: 40
        },
        existing: [
          {
            id: 100,
            teacher_id: 10,
            school_id: 1,
            group_id: 2,
            room_id: 6,
            series_id: 2,
            day: 1,
            start_slot: 4,
            end_slot: 5,
            enrollment: 20,
            room_capacity: 40
          }
        ],
        state: 'BORRADOR'
      };
      const res = await fetch('/api/v1/validaciones/placement',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(payload)});
      if(!res.ok){
        const errorBody = await res.json().catch(()=>({error:res.statusText}));
        throw new Error(errorBody.error || 'HTTP '+res.status);
      }
      const json = await res.json(); el('#valResult').textContent = JSON.stringify(json,null,2); setStatus('validación completa');
    }catch(e){el('#valResult').textContent = 'Error: '+e.message; setStatus('error en validación')}
  })
  el('#runComplex').addEventListener('click', async ()=>{
    setStatus('ejecutando validación compleja');
    try{
      const payload = {
        proposed: {
          id: 201,
          teacher_id: 10,
          school_id: 1,
          group_id: 1,
          course_id: 101,
          series_id: 1,
          room_id: 5,
          pavilion_id: 2,
          day: 1,
          start_slot: 3,
          end_slot: 4,
          enrollment: 55,
          room_capacity: 40
        },
        existing: [
          {
            id: 200,
            teacher_id: 10,
            school_id: 1,
            group_id: 2,
            course_id: 102,
            series_id: 2,
            room_id: 6,
            pavilion_id: 1,
            day: 1,
            start_slot: 4,
            end_slot: 5,
            enrollment: 20,
            room_capacity: 40
          }
        ],
        distances: [
          {from_pavilion_id: 1, to_pavilion_id: 2, minutes: 90}
        ],
        department_sessions: [
          {department_id: 7, day: 1, start_slot: 4, end_slot: 5}
        ],
        state: 'EN_REAJUSTE'
      };
      const res = await fetch('/api/v1/validaciones/placement',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(payload)});
      if(!res.ok){
        const errorBody = await res.json().catch(()=>({error:res.statusText}));
        throw new Error(errorBody.error || 'HTTP '+res.status);
      }
      const json = await res.json(); el('#valResult').textContent = JSON.stringify(json,null,2); setStatus('validación completa');
    }catch(e){el('#valResult').textContent = 'Error: '+e.message; setStatus('error en validación')}
  })
}

async function loadValidationScenarios(){
  const result = el('#scenarioResult');
  if(!result) return;
  result.textContent = 'Cargando escenarios del motor...';
  try{
    const scenarios = await apiGET('/validaciones/escenarios');
    result.textContent = JSON.stringify(scenarios, null, 2);
  }catch(e){
    result.textContent = 'Error al cargar escenarios: '+e.message;
  }
}

async function renderSchedules(container){
  setStatus('mostrando horarios');
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Horarios</h3><p class="small">Vista rápida de horarios existentes (de prueba).</p><div id="schedulesList"></div>`;
  container.appendChild(card);
  try{
    const hs = await apiGET('/horarios');
    const pre = document.createElement('pre'); pre.className='preview'; pre.textContent = JSON.stringify(hs,null,2); el('#schedulesList').appendChild(pre);
    setStatus('horarios cargados');
  }catch(e){el('#schedulesList').textContent='Error: '+e.message; setStatus('error al cargar horarios')}
}

function renderSiige(container){
  setStatus('mostrando simulador SIIGE');
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Simulador SIIGE</h3><p class="small">Pegar aquí un JSON simulado de SIIGE para probar importación (local).</p><textarea id="siigeJson" style="width:100%;height:200px"></textarea><div class="controls"><button type="button" class="btn" id="simulate">Simular Import</button></div><pre id="siigeResult" class="preview"></pre>`;
  container.appendChild(card);
  el('#simulate').addEventListener('click', ()=>{
    const txt = el('#siigeJson').value;
    try{
      const parsed = JSON.parse(txt);
      el('#siigeResult').textContent = JSON.stringify({ok:true,importPreview:parsed},null,2);
      setStatus('simulación lista');
    }catch(e){el('#siigeResult').textContent='JSON inválido: '+e.message; setStatus('JSON inválido')}
  })
}

// setup nav: prefer event delegation if buttons are not found later
const navButtons = els('.nav-btn');
if(navButtons.length>0){
  navButtons.forEach(b=>b.addEventListener('click',e=>{ const r = b.getAttribute('data-route'); render(r); }));
} else {
  document.addEventListener('click', e=>{
    const btn = e.target.closest && e.target.closest('.nav-btn');
    if(btn){ render(btn.getAttribute('data-route')) }
  })
}

render('dashboard');
