import React from 'react'
import ReactDOM from 'react-dom/client'
import LogRocket from 'logrocket'
import { TabProvider } from './Context/Tabs'
import ErrorBoundary from './Components/ErrorBoundary'
import './index.css'
import App from './App'

import reportWebVitals from './reportWebVitals'
import * as serviceWorkerRegistration from './serviceWorkerRegistration'

import { initializeBridge } from 'birdcash-miniprogram-sdk-alpha'

LogRocket.init('3ouaid/miniprogram-developer')

initializeBridge()
  .then(() => {
    const root = ReactDOM.createRoot(
      document.getElementById('root') as HTMLElement,
    )

    root.render(
      <React.StrictMode>
        <ErrorBoundary>
          <TabProvider>
            <App />
          </TabProvider>
        </ErrorBoundary>
      </React.StrictMode>,
    )

    // If you want your app to work offline and load faster, you can change
    // unregister() to register() below. Note this comes with some pitfalls.
    // Learn more about service workers: https://cra.link/PWA
    serviceWorkerRegistration.unregister()
    //serviceWorkerRegistration.register()

    // If you want to start measuring performance in your app, pass a function
    // to log results (for example: reportWebVitals(console.log))
    // or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
    reportWebVitals()
  })
  .catch(console.error)

/*
const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)

root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorkerRegistration.unregister()
//serviceWorkerRegistration.register()

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
*/
