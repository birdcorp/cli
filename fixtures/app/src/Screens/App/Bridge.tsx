import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function Bridge() {
  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Initialize Bridge',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 mt-5 pb-10">
      <h3 className="text-white font-bold mb-3">Initalize Bridge</h3>
      <p className="text-white mt-10">For typescript support use npm</p>

      <p className="text-white mt-10 font-bold">ESM</p>

      <p className="text-white my-8">Load the script in the head tag</p>

      <HighlightedCode
        code={`<script type="module">
  import { initializeBridge } from "https://unpkg.com/birdcash-miniprogram-sdk-alpha@latest";

  try {
    await initializeBridge();
  } catch (error) {
    console.error("Bridge initialization failed:", error);
  }
</script>`}
      />

      <p className="text-white mt-10 font-bold mb-5">Npm</p>
      <HighlightedCode
        className="mb-5"
        code={`npm install birdcash-miniprogram-sdk-alpha`}
      />

      <p className="text-white mt-10">
        Initialize the bridge first then render the ui
      </p>
      <p className="text-white mt-8 mb-6">Example: React</p>

      <HighlightedCode
        code={`import { initializeBridge } from "birdcash-miniprogram-sdk-alpha";

initializeBridge()
  .then(() => {
    const root = ReactDOM.createRoot(
      document.getElementById('root') as HTMLElement,
    )

    root.render(
      <React.StrictMode>
        <RouterProvider router={router} />
      </React.StrictMode>,
    )
  })
  .catch(console.error)`}
      />

      <p className="text-white mt-10 mb-6">
        All set! you can now call the native code using
      </p>

      <HighlightedCode
        className="mb-5"
        code={`window.miniProgram?.call('exampleHandler', {})`}
      />
    </div>
  )
}
