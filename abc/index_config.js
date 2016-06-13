import menu from './index_menu';
export default function config(NgAdminConfigurationProvider) {
  let nga = NgAdminConfigurationProvider;
  let admin = nga.application('Dashboard');
  //.baseApiUrl('http://localhost:9090/api/v1/');
  // let workloads = nga.entity('workloads');
  /* workloads.creationView().fields([
     nga.field('podList.pods.objectMeta.name').label('PodsName'),
     nga.field('podList.pods.podPhase').label('PodPhase'),
   ]);
 */
  // workloads.editionView().fields(workloads.creationView().fields());
  // admin.addEntity(workloads);
  admin.menu(menu(nga));
  nga.configure(admin);
}
