package validationui

const pageHTML = `<!doctype html>
<html lang="es">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Validaciones de Horarios UNSCH</title>
  <style>
    :root {
      color-scheme: light;
      --ink: #1c2430;
      --muted: #5f6f82;
      --line: #d9e0e8;
      --surface: #ffffff;
      --bg: #f4f7f5;
      --blocker: #b42318;
      --warning: #a15c07;
      --info: #1d5f95;
      --ok: #177245;
    }
    * { box-sizing: border-box; }
    body { margin: 0; font-family: Arial, Helvetica, sans-serif; background: var(--bg); color: var(--ink); }
    header { padding: 28px 32px 18px; border-bottom: 1px solid var(--line); background: #fff; }
    h1 { margin: 0 0 8px; font-size: 28px; font-weight: 700; }
    p { margin: 0; color: var(--muted); line-height: 1.45; }
    main { padding: 24px 32px 40px; max-width: 1180px; margin: 0 auto; }
    .toolbar { display: flex; gap: 12px; align-items: center; justify-content: space-between; margin-bottom: 18px; flex-wrap: wrap; }
    button { border: 1px solid #23374d; background: #23374d; color: white; height: 38px; padding: 0 14px; border-radius: 6px; cursor: pointer; font-weight: 700; }
    button:hover { background: #172637; }
    .summary { display: flex; gap: 10px; flex-wrap: wrap; }
    .pill { border: 1px solid var(--line); background: #fff; padding: 8px 10px; border-radius: 999px; font-size: 13px; font-weight: 700; }
    .grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr)); gap: 14px; }
    .card { background: var(--surface); border: 1px solid var(--line); border-radius: 8px; padding: 16px; min-height: 190px; display: flex; flex-direction: column; gap: 12px; }
    .card header { padding: 0; border: 0; background: transparent; display: flex; justify-content: space-between; gap: 12px; align-items: flex-start; }
    .title { font-size: 17px; font-weight: 700; }
    .rule { font-size: 12px; font-weight: 800; border: 1px solid var(--line); border-radius: 999px; padding: 5px 8px; white-space: nowrap; }
    .desc { color: var(--muted); font-size: 14px; }
    .panel-grid { display: grid; grid-template-columns: 2fr 1fr; gap: 16px; align-items: start; }
    .panel { background: var(--surface); border: 1px solid var(--line); border-radius: 8px; padding: 16px; display: flex; flex-direction: column; gap: 12px; }
    .controls { display: flex; gap: 10px; flex-wrap: wrap; }
    .secondary { background: #fff; color: #23374d; border-color: var(--line); }
    .secondary:hover { background: #f3f6f9; }
    textarea { width: 100%; min-height: 240px; border: 1px solid var(--line); border-radius: 8px; padding: 12px; font-family: Consolas, monospace; font-size: 13px; resize: vertical; }
    .preview { background: #f7fafc; border: 1px solid var(--line); border-radius: 8px; padding: 12px; white-space: pre-wrap; word-break: break-word; font-family: Consolas, monospace; font-size: 13px; }
    .finding { border-left: 4px solid var(--line); padding: 10px 10px 10px 12px; background: #f9fbfc; border-radius: 4px; font-size: 14px; }
    .finding strong { display: block; margin-bottom: 4px; }
    .BLOCKER { border-left-color: var(--blocker); }
    .WARNING { border-left-color: var(--warning); }
    .INFO { border-left-color: var(--info); }
    .OK { border-left-color: var(--ok); }
    .badge { font-size: 12px; font-weight: 800; padding: 5px 8px; border-radius: 999px; color: #fff; }
    .badge.BLOCKER { background: var(--blocker); }
    .badge.WARNING { background: var(--warning); }
    .badge.INFO { background: var(--info); }
    .badge.OK { background: var(--ok); }
    code { background: #edf2f7; padding: 2px 5px; border-radius: 4px; }
    @media (max-width: 920px) { .panel-grid { grid-template-columns: 1fr; } }
    @media (max-width: 680px) { header, main { padding-left: 16px; padding-right: 16px; } h1 { font-size: 23px; } }
  </style>
</head>
<body>
  <header>
    <h1>Validaciones de Horarios UNSCH</h1>
    <p>Panel visual para comprobar reglas de negocio del motor de horarios. Los resultados salen del backend Go, no de datos ficticios en el navegador.</p>
  </header>
  <main>
    <div class="toolbar">
      <button id="run">Ejecutar validaciones</button>
      <div class="summary" id="summary"></div>
    </div>
    <div class="panel-grid">
      <section class="panel">
        <h2 style="margin: 0; font-size: 20px;">Escenarios del motor</h2>
        <p class="desc">Cada tarjeta representa un caso de prueba alimentado por el backend Go.</p>
        <section class="grid" id="grid"></section>
      </section>
      <section class="panel">
        <h2 style="margin: 0; font-size: 20px;">Validar payload personalizado</h2>
        <p class="desc">Pega un JSON con el formato esperado por el endpoint de placement y comprueba si el motor lo acepta o lo bloquea.</p>
        <textarea id="customPayload" spellcheck="false"></textarea>
        <div class="controls">
          <button id="runCustom" type="button">Validar payload</button>
          <button id="loadSample" class="secondary" type="button">Cargar ejemplo</button>
        </div>
        <pre id="customResult" class="preview">Esperando validación...</pre>
      </section>
    </div>
  </main>
  <script>
    const grid = document.querySelector('#grid');
    const summary = document.querySelector('#summary');
    const run = document.querySelector('#run');
    const customPayload = document.querySelector('#customPayload');
    const customResult = document.querySelector('#customResult');
    const runCustom = document.querySelector('#runCustom');
    const loadSample = document.querySelector('#loadSample');

    function severityLabel(severity) {
      if (severity === 'BLOCKER') return 'Bloqueo';
      if (severity === 'WARNING') return 'Advertencia';
      if (severity === 'INFO') return 'Informativo';
      return 'Valido';
    }

    function escapeHtml(value) {
      return String(value)
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/\"/g, '&quot;');
    }

    function renderSummary(items) {
      const counts = { BLOCKER: 0, WARNING: 0, INFO: 0, OK: 0 };
      for (const item of items) {
        if (!item.findings || item.findings.length === 0) counts.OK++;
        for (const finding of item.findings || []) counts[finding.severity]++;
      }
      summary.innerHTML = '<span class="pill">Bloqueos: ' + counts.BLOCKER + '</span>' +
        '<span class="pill">Advertencias: ' + counts.WARNING + '</span>' +
        '<span class="pill">Informativos: ' + counts.INFO + '</span>' +
        '<span class="pill">Validos: ' + counts.OK + '</span>';
    }

    function render(items) {
      renderSummary(items);
      grid.innerHTML = items.map(function(item) {
        const findings = item.findings && item.findings.length ? item.findings : [{ rule: item.rule, severity: 'OK', message: 'La asignacion no genera observaciones.' }];
        const findingHtml = findings.map(function(finding) {
          return '<div class="finding ' + finding.severity + '">' +
            '<strong><span class="badge ' + finding.severity + '">' + severityLabel(finding.severity) + '</span></strong>' +
            '<span><code>' + escapeHtml(finding.rule) + '</code> ' + escapeHtml(finding.message) + '</span>' +
            '</div>';
        }).join('');
        return '<article class="card">' +
          '<header>' +
            '<div class="title">' + escapeHtml(item.title) + '</div>' +
            '<span class="rule">' + escapeHtml(item.rule) + '</span>' +
          '</header>' +
          '<p class="desc">' + escapeHtml(item.description) + '</p>' +
          findingHtml +
          '</article>';
      }).join('');
    }

    function defaultPayload() {
      return JSON.stringify({
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
        existing: [{
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
        }],
        state: 'BORRADOR'
      }, null, 2);
    }

    async function load() {
      run.disabled = true;
      run.textContent = 'Ejecutando...';
      try {
        const response = await fetch('/api/v1/validaciones/escenarios');
        if (!response.ok) throw new Error('HTTP ' + response.status);
        render(await response.json());
      } catch (error) {
        grid.innerHTML = '<article class="card"><div class="finding BLOCKER"><strong>Error</strong>' + escapeHtml(error.message) + '</div></article>';
      } finally {
        run.disabled = false;
        run.textContent = 'Ejecutar validaciones';
      }
    }

    async function validateCustom() {
      runCustom.disabled = true;
      runCustom.textContent = 'Validando...';
      customResult.textContent = 'Enviando payload...';
      try {
        const payload = JSON.parse(customPayload.value);
        const response = await fetch('/api/v1/validaciones/placement', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload)
        });
        if (!response.ok) throw new Error('HTTP ' + response.status);
        const findings = await response.json();
        customResult.textContent = JSON.stringify(findings, null, 2);
      } catch (error) {
        customResult.textContent = 'Error: ' + error.message;
      } finally {
        runCustom.disabled = false;
        runCustom.textContent = 'Validar payload';
      }
    }

    loadSample.addEventListener('click', function() {
      customPayload.value = defaultPayload();
    });

    runCustom.addEventListener('click', validateCustom);
    run.addEventListener('click', load);
    customPayload.value = defaultPayload();
    load();
  </script>
</body>
</html>`
