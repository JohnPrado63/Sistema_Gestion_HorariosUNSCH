const API_BASE = '/api/v1';

function el(q, parent=document) { return parent.querySelector(q); }
function els(q, parent=document) { return Array.from(parent.querySelectorAll(q)); }

async function apiGET(path){
  const res = await fetch(API_BASE + path);
  if(!res.ok) throw new Error('API error '+res.status);
  return res.json();
}

function render(route){
  const main = el('#main');
  main.innerHTML = '';
  if(route==='dashboard') return renderDashboard(main);
  if(route==='catalogs') return renderCatalogs(main);
  if(route==='validations') return renderValidations(main);
  if(route==='schedules') return renderSchedules(main);
  if(route==='siige') return renderSiige(main);
  renderDashboard(main);
}

async function renderDashboard(container){
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Resumen</h3><p class="small">Demo local del Motor de Horarios — APIs disponibles: <span class="code">/api/v1</span></p><div class="grid"><div class="card"><h4>Validaciones</h4><p class="small">RV-01..RV-09 disponibles en el backend.</p></div><div class="card"><h4>Importador SIIGE</h4><p class="small">Simulador en esta SPA para probar importes.</p></div></div>`;
  container.appendChild(card);
}

async function renderCatalogs(container){
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Catálogos</h3><p class="small">Cargar listados desde API para revisión.</p><div class="controls"><button class="btn" id="loadCatalogs">Cargar</button></div><div id="catalogResult" class="preview"></div>`;
  container.appendChild(card);
  el('#loadCatalogs').addEventListener('click',async ()=>{
    try{
      const facs = await apiGET('/facultades');
      const esc = await apiGET('/escuelas');
      document.getElementById('catalogResult').textContent = JSON.stringify({facultades:facs,escuelas:esc},null,2);
    }catch(e){document.getElementById('catalogResult').textContent = 'Error: '+e.message}
  })
}

async function renderValidations(container){
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Validaciones</h3><p class="small">Enviar escenarios de prueba al backend.</p><div class="controls"><button class="btn" id="runSimple">Escenario simple</button><button class="btn secondary" id="runComplex">Escenario complejo</button></div><pre id="valResult" class="preview"></pre>`;
  container.appendChild(card);
  el('#runSimple').addEventListener('click', async ()=>{
    try{
      const payload = {horarios:[],grupos:[],docentes:[],periodo:'2026-I'};
      const res = await fetch('/api/v1/validaciones/placement',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(payload)});
      const json = await res.json(); el('#valResult').textContent = JSON.stringify(json,null,2);
    }catch(e){el('#valResult').textContent = 'Error: '+e.message}
  })
  el('#runComplex').addEventListener('click', async ()=>{
    try{
      // sample complex scenario (local demo)
      const payload = {periodo:'2026-I',horarios:[{curso:'MAT101',aula:'A101',dia:1,inicio:'08:00',fin:'10:00',docente:'d1'},{curso:'FIS101',aula:'A102',dia:1,inicio:'08:30',fin:'10:30',docente:'d1'}],grupos:[],docentes:[{id:'d1',nombre:'Juan'}]};
      const res = await fetch('/api/v1/validaciones/placement',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(payload)});
      const json = await res.json(); el('#valResult').textContent = JSON.stringify(json,null,2);
    }catch(e){el('#valResult').textContent = 'Error: '+e.message}
  })
}

async function renderSchedules(container){
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Horarios</h3><p class="small">Vista rápida de horarios existentes (de prueba).</p><div id="schedulesList"></div>`;
  container.appendChild(card);
  try{
    const hs = await apiGET('/horarios');
    const pre = document.createElement('pre'); pre.className='preview'; pre.textContent = JSON.stringify(hs,null,2); el('#schedulesList').appendChild(pre);
  }catch(e){el('#schedulesList').textContent='Error: '+e.message}
}

function renderSiige(container){
  const card = document.createElement('div'); card.className='card';
  card.innerHTML = `<h3>Simulador SIIGE</h3><p class="small">Pegar aquí un JSON simulado de SIIGE para probar importación (local).</p><textarea id="siigeJson" style="width:100%;height:200px"></textarea><div class="controls"><button class="btn" id="simulate">Simular Import</button></div><pre id="siigeResult" class="preview"></pre>`;
  container.appendChild(card);
  el('#simulate').addEventListener('click', ()=>{
    const txt = el('#siigeJson').value;
    try{
      const parsed = JSON.parse(txt);
      el('#siigeResult').textContent = JSON.stringify({ok:true,importPreview:parsed},null,2);
    }catch(e){el('#siigeResult').textContent='JSON inválido: '+e.message}
  })
}

// setup nav
els('.nav-btn').forEach(b=>b.addEventListener('click',e=>render(b.getAttribute('data-route'))));
render('dashboard');
