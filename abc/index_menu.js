export default function menu(nga) {
  return nga.menu().addChild(nga.menu()
                                 .title('OverView')
                                 .icon('<span class="fa fa-dashboard"></span>')
                                 .active(path => path.indexOf('/workloads') === 0)
                                 .addChild(nga.menu()
                                               .title('Services')
                                               .link('/services')
                                               .icon('<span class="fa fa-cubes"></span> '))
                                 .addChild(nga.menu()
                                               .title('RC')
                                               .link('/replicationcontrollers')
                                               .icon('<span class="fa fa-life-bouy"></span>'))
                                 .addChild(nga.menu().title('Pods').link('/pods').icon(
                                     '<span class="fa fa-cube"></span>')))
                            .addChild(nga.menu()
                            .title('Docker Images')
                            .icon('<span class="fa fa-cloud"></span>'));
}
