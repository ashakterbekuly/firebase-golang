// const headerTemplate = document.createElement('template');
// headerTemplate.innerHTML = `

// <style>
//     .header{
//         text-align: center;
//     }
//     h1{
//         color: blue;
//     }
// </style>

// <nav id="header" class="fixed w-full z-30 top-0 text-white">
//       <div class="w-full container mx-auto flex flex-wrap items-center justify-between mt-0 py-2">
//         <div class="pl-4 flex items-center align-middle">
//           <a class="toggleColour align-middle text-black no-underline hover:no-underline font-bold lg:text-4xl inline-block" href="#">
//             <!-- Icon from: http://www.potlabicons.com/ -->
//             <img src="/img/logo_archneo_black.png" alt="" class="w-12 inline-block align-middle">
//             <div class="h-22 inline-block">ArchNEO</div>
//           </a>
//         </div>
//         <div class="block lg:hidden pr-4">
//           <button id="nav-toggle" class="flex items-center p-1 text-pink-800 hover:text-gray-900 focus:outline-none focus:shadow-outline transform transition hover:scale-105 duration-300 ease-in-out">
//             <svg class="fill-current h-6 w-6" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
//               <title>Menu</title>
//               <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
//             </svg>
//           </button>
//         </div>
//         <div class="w-full flex-grow lg:flex lg:items-center lg:w-auto hidden mt-2 lg:mt-0 bg-white lg:bg-transparent text-black p-4 lg:p-0 z-20" id="nav-content">
//           <ul class="list-reset lg:flex justify-end flex-1 items-center">
//             <li class="mr-3">
//               <a class="inline-block py-2 px-4 text-black font-bold text-lg no-underline" href="/">Home</a>
//             </li>
//             <li class="mr-3">
//               <a class="inline-block text-black text-lg no-underline hover:text-gray-800 hover:text-underline py-2 px-4" href="/events">Events</a>
//             </li>
//             <li class="mr-3">
//               <a class="inline-block text-black text-lg no-underline hover:text-gray-800 hover:text-underline py-2 px-4" href="/projects">Projects</a>
//             </li>
//             <li class="mr-3">
//               <a href="/register_role.html">
//                 <button class="lg:mx-0 hover:underline bg-gray-800 text-xl text-white rounded-full my-4 py-2 px-8 shadow-lg focus:outline-none focus:shadow-outline transform transition hover:scale-105 duration-300 ease-in-out">
//                   Log In
//                 </button>
//               </a>
//             </li>
//             <li class="mr-3">
//               <a href="/register_role.html">
//                 <button class="lg:mx-0 hover:underline bg-gray-200 text-xl text-gray-800 rounded-full my-4 py-2 px-8 shadow-lg focus:outline-none focus:shadow-outline transform transition hover:scale-105 duration-300 ease-in-out">
//                   Register
//                 </button>
//               </a>
//             </li>
//           </ul>
//           <!-- <button
//             id="navAction"
//             class="mx-auto lg:mx-0 hover:underline bg-white text-gray-800 font-bold rounded-full mt-4 lg:mt-0 py-4 px-8 shadow opacity-75 focus:outline-none focus:shadow-outline transform transition hover:scale-105 duration-300 ease-in-out"
//           >
//             Register
//           </button>
//           <a href="/login.html">
//             <button id="navAction" class="mx-auto lg:mx-0 hover:underline bg-white text-gray-800 font-bold rounded-full my-6 py-4 px-12 shadow-lg focus:outline-none focus:shadow-outline transform transition hover:scale-105 duration-300 ease-in-out">
//               Login
//             </button> -->
//             </a>
//         </div>
//       </div>
//       <hr class="border-b border-gray-100 opacity-25 my-0 py-0" />
//     </nav>
// `

class Header extends HTMLElement {
    constructor() {
        // Always call super first in constructor
        super();
        const shadow = this.attachShadow({ mode: 'open' });

        const nav = document.createElement('nav');
        nav.classList.add('fixed w-full z-30 top-0 text-white');

        const wrapper = document.createElement('div');
        wrapper.classList.add('w-full container mx-auto flex flex-wrap items-center justify-between mt-0 py-2');
        nav.appendChild(wrapper);

        const container = document.createElement('div');
        container.classList.add('pl-4 flex items-center align-middle');
        wrapper.appendChild(container);

        const link = document.createElement('a');
        link.classList.add('toggleColour align-middle text-black no-underline hover:no-underline font-bold lg:text-4xl inline-block');
        container.appendChild(link);

        const img = document.createElement('img');
        img.setAttribute = ('src', '/img/logo_archneo_black.png');
        img.classList.add('w-12 inline-block align-middle');
        link.appendChild(img);

        shadow.appendChild(nav);
    }
    

    // connectedCallback() {
    //     const shadowRoot = this.attachShadow({ mode: 'open' });
    //     shadowRoot.appendChild(headerTemplate.content);
    // }
}

customElements.define('header-component', Header);