const modal = document.querySelector('.modal');
            const showModal = document.querySelectorAll('.more-info');
            const closeModal = document.querySelectorAll('.close-modal');
            const register = document.querySelector('register');
    
            showModal.forEach( show => {
              show.addEventListener('click', function(){
                  modal.classList.remove('hidden')
              });
            });
            
    
            closeModal.forEach( close => {
              close.addEventListener('click', function(){
                modal.classList.add('hidden')
              });
            });
            
            register.addEventListener('click', function(){
                modal.classList.add('hidden');
    
            });