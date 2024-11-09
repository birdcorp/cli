import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function Alert() {
  const code = `window.miniProgram?.call("showAlert", {
    title: "Attention",
    message: "This is an alert",
    success: (response) => {
        if(response.ok){
            // confirm button was pressed
        }
    }
})`

  function trigger() {
    console.log('trigger alert')
    window.miniProgram?.call(
      'showAlert',
      {
        title: 'Attention',
        message: 'This is an alert',
      },
      {
        success(response: any) {
          console.log('success..')
          if (response.ok) {
            // confirm button was pressed
            window.miniProgram?.call('showToast', {
              type: 'success',
              content: 'OK Pressed',
              duration: 3000,
            })
          } else {
            window.miniProgram?.call(
              'showToast',
              {
                type: 'success',
                content: 'Cancel Pressed',
                duration: 3000,
              },
              {
                success(response: any) {
                  if (response.ok) {
                    // confirm button was pressed
                  }
                },
                failure(err: Error) {
                  console.error(err)
                },
                completed() {},
              },
            )
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  useEffect(() => {
    async function onMount() {
      window.miniProgram?.call(
        'setNavigationBar',
        {
          title: 'Alert',
          backgroundColor: '#0A0A0A',
          color: 'light',
        },
        {
          success(response: any) {
            if (response.ok) {
              // confirm button was pressed
            }
          },
          failure(err: any) {
            console.error(err)
          },
          completed() {},
        },
      )

      window.alert = function (message) {
        window.miniProgram?.call('showAlert', {
          title: 'Alert',
          message,
        })
      }
    }

    onMount()
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Alert</h3>

      <p className="text-gray-300 mb-8">
        Triggers a native alert with a title and message properties
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Show
      </button>

      <h3 className="text-white mt-8 font-bold">Code</h3>
      <HighlightedCode code={code} />

      <button onClick={() => window.alert('this is a monkey patched alert')}>
        Native alert
      </button>
    </div>
  )
}
