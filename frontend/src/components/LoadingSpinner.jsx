import React from 'react';

const LoadingSpinner = ({ message = "Running tests..." }) => {
  return (
    <div className="text-center py-12">
      <div className="relative inline-block">
        {/* Outer rotating ring */}
        <div className="w-16 h-16 border-4 border-dark-border rounded-full animate-spin">
          <div className="w-full h-full border-4 border-transparent border-t-neon-orange rounded-full animate-pulse"></div>
        </div>
        
        {/* Inner pulsing dot */}
        <div className="absolute inset-0 flex items-center justify-center">
          <div className="w-3 h-3 bg-neon-orange rounded-full animate-ping"></div>
        </div>
        
        {/* Neon glow effect */}
        <div className="absolute inset-0 w-16 h-16 border-2 border-neon-orange/30 rounded-full animate-pulse shadow-neon"></div>
      </div>
      
      <div className="mt-6 space-y-2">
        <p className="text-lg font-semibold text-neon-orange animate-pulse">
          {message}
        </p>
        
        <div className="flex items-center justify-center space-x-1">
          <div className="w-2 h-2 bg-neon-orange rounded-full animate-bounce"></div>
          <div className="w-2 h-2 bg-neon-orange rounded-full animate-bounce" style={{ animationDelay: '0.1s' }}></div>
          <div className="w-2 h-2 bg-neon-orange rounded-full animate-bounce" style={{ animationDelay: '0.2s' }}></div>
        </div>
        
        <p className="text-sm text-dark-text-secondary">
          Making API requests and validating responses...
        </p>
      </div>
    </div>
  );
};

export default LoadingSpinner;