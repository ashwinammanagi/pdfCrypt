<script lang="ts">
  import { PickPDF, EncryptPDF, DecryptPDF } from "../wailsjs/go/main/App.js";

  let mode: "encrypt" | "decrypt" = "encrypt";
  let inputPath = "";
  let password = "";
  let resultMessage = "";

  $: toggleLabel = permissions.every((p) => !p.checked)
    ? "Select All"
    : "Unselect All";

  const allPermissions = [
    { name: "Print", value: 1 },
    { name: "Modify", value: 2 },
    { name: "Extract", value: 4 },
    { name: "Annotate", value: 8 },
    { name: "Fill Forms", value: 16 },
    { name: "Assemble", value: 32 },
    { name: "Print High Resolution", value: 64 },
  ];
  let permissions = allPermissions.map((p) => ({ ...p, checked: true }));

  async function chooseFile() {
    try {
      const path = await PickPDF();
      inputPath = path || "No file selected.";
    } catch (err) {
      inputPath = "Error: " + err;
    }
  }

  function togglePermissions() {
    const allUnchecked = permissions.every((p) => !p.checked);
    permissions = permissions.map((p) => ({ ...p, checked: allUnchecked }));
  }

  function processPDF() {
    if (!inputPath || !password) {
      resultMessage = "Please select a file and enter a password.";
      return;
    }

    const selectedPermissions = permissions
      .filter((p) => p.checked)
      .reduce((acc, p) => acc | p.value, 0);

    const action =
      mode === "encrypt"
        ? () => EncryptPDF(inputPath, password, selectedPermissions)
        : () => DecryptPDF(inputPath, password);

    action()
      .then((output) => {
        resultMessage = `${mode === "encrypt" ? "Encrypted" : "Decrypted"} file saved to: ${output}`;
      })
      .catch((err) => {
        resultMessage = `Error: ${err}`;
      });
  }
</script>

<main>
  <h1 class="title">pdfCrypt</h1>

  <div class="mode-toggle">
    <label
      ><input type="radio" bind:group={mode} value="encrypt" /> Encrypt</label
    >
    <label
      ><input type="radio" bind:group={mode} value="decrypt" /> Decrypt</label
    >
  </div>

  <div class="form-group">
    <button class="btn" on:click={chooseFile}>Choose PDF</button>
    <div class="file-path">{inputPath}</div>

    <input
      class="input"
      placeholder="Enter password"
      type="password"
      bind:value={password}
    />

    {#if mode === "encrypt"}
      <fieldset class="permissions">
        <legend>Permissions</legend>
        {#each permissions as perm}
          <label class="perm">
            <input type="checkbox" bind:checked={perm.checked} />
            {perm.name}
          </label>
        {/each}
        <button type="button" class="btn small" on:click={togglePermissions}>
          {toggleLabel}
        </button>
      </fieldset>
    {/if}

    <button class="btn primary" on:click={processPDF}>
      {mode === "encrypt" ? "Encrypt" : "Decrypt"} PDF
    </button>
  </div>

  {#if resultMessage}
    <p class="result">{resultMessage}</p>
  {/if}
</main>

<style>
  :global(body),
  :global(html) {
    margin: 0;
    padding: 0;
    height: 100%;
    background-color: #0f1a23;
    overflow-x: hidden;
    overflow-y: hidden;
  }

  main {
    max-width: 420px;
    margin: 5vh auto;
    padding: 2rem 1rem;
    font-family: system-ui, sans-serif;
    box-sizing: border-box;
  }

  .title {
    font-size: 1.8rem;
    margin-bottom: 1rem;
    color: #ffffff;
    text-align: center;
  }

  .mode-toggle {
    display: flex;
    justify-content: center;
    gap: 1.5rem;
    margin-bottom: 1rem;
    font-size: 0.95rem;
    color: #cfd8dc;
  }

  .form-group {
    display: flex;
    flex-direction: column;
    gap: 0.9rem;
    width: 100%;
  }

  .btn,
  .input,
  .file-path,
  .permissions {
    width: 100%;
    box-sizing: border-box;
  }

  .btn {
    background-color: #394b59;
    color: white;
    border: none;
    padding: 0.6rem 1rem;
    border-radius: 0.5rem;
    cursor: pointer;
    font-size: 0.95rem;
    transition: background-color 0.2s ease;
  }

  .btn:hover {
    background-color: #526170;
  }

  .btn.primary {
    background-color: #2e8b57;
  }

  .btn.primary:hover {
    background-color: #3aa06e;
  }

  .file-path {
    font-size: 0.8rem;
    color: #9faebf;
    word-break: break-word;
  }

  .input {
    padding: 0.6rem 1rem;
    border-radius: 0.5rem;
    border: 1px solid #4b5a6a;
    background-color: #2a3b4c;
    color: #ffffff;
    font-size: 0.95rem;
  }

  .input::placeholder {
    color: #8799a7;
  }

  fieldset.permissions {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 0.3rem 1rem;
    padding: 0.8rem 1rem;
    border: 1px solid #4b5a6a;
    border-radius: 0.5rem;
    background-color: #1f2d3a;
    color: #cfd8dc;
    font-size: 0.85rem;
    width: 100%;
    box-sizing: border-box;
  }

  fieldset.permissions legend {
    font-size: 0.9rem;
    margin-bottom: 0.3rem;
    padding: 0 0.5rem;
    color: #cfd8dc;
  }

  .perm {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .result {
    margin-top: 1.5rem;
    font-size: 0.9rem;
    color: #b2f5ea;
    word-break: break-word;
    text-align: center;
  }

  .btn.small {
    padding: 0.1rem 0.5rem;
    font-size: 0.9rem;
    background-color: #556c7a;
  }

  .btn.small:hover {
    background-color: #6d8493;
  }
</style>
