import React from 'react'
import '../App.css'; 
 

export default function Header() {
  return (
    <header className="bg-dark-surface border-b border-dark-border shadow-lg ">
    <div className="max-w-7xl mx-auto px-4 py-6">
      <div className="flex items-center space-x-4">
      
        <div>
          <h1 className="text-3xl font-bold text-neon-orange neon-text header-logo" >
            COMAPI
          </h1>
          <p className="text-dark-text-secondary text-sm">
            Navigate your API testing with confidence
          </p>
        </div>
      </div>
    </div>
  </header>
  )
}
